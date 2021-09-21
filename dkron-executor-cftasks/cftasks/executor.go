package cftasks

import (
	"fmt"
	"log"
	"time"

	"code.cloudfoundry.org/bytefmt"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	"github.com/armon/circbuf"
	"github.com/distribworks/dkron/v3/plugin"
	"github.com/distribworks/dkron/v3/plugin/types"
	"github.com/orange-cloudfoundry/cfron/dkron-executor-cftasks/configs"
	"github.com/orange-cloudfoundry/cfron/dkron-executor-cftasks/sessions"
)

const (
	// maxBufSize limits how much data we collect from a handler.
	// This is to prevent Serf's memory from growing to an enormous
	// amount due to a faulty handler.
	maxBufSize = 256000
)

type reportingWriter struct {
	buffer  *circbuf.Buffer
	cb      plugin.StatusHelper
	isError bool
}

func (p reportingWriter) Write(data []byte) (n int, err error) {
	p.cb.Update(data, false)
	return p.buffer.Write(data)
}

type Executor struct {
}

func (e Executor) Execute(args *types.ExecuteRequest, cb plugin.StatusHelper) (*types.ExecuteResponse, error) {
	output, _ := circbuf.NewBuffer(maxBufSize)

	config, err := configs.GetConfig()
	if err != nil {
		return nil, err
	}

	sess, err := sessions.GetSession()
	if err != nil {
		return nil, err
	}

	client := sess.V3()

	command, ok := args.Config["command"]
	if !ok {
		return nil, fmt.Errorf("command must be given")
	}

	appGUID, ok := args.Config["app_guid"]
	if !ok {
		return nil, fmt.Errorf("app_guid must be given")
	}

	diskInMB := uint64(256)
	memoryInMB := uint64(64)

	diskUserSet := false
	diskRaw, ok := args.Config["disk"]
	if ok && diskRaw != "" {
		diskInMB, err = bytefmt.ToMegabytes(diskRaw)
		if err != nil {
			return nil, err
		}
		diskUserSet = true
	}

	memUserSet := false
	memRaw, ok := args.Config["memory"]
	if ok && memRaw != "" {
		memoryInMB, err = bytefmt.ToMegabytes(memRaw)
		if err != nil {
			return nil, err
		}
		memUserSet = true
	}

	if !diskUserSet || !memUserSet {
		process, _, err := client.GetApplicationProcessByType(appGUID, constant.ProcessTypeWeb)
		// we took default value from process web if exists
		if err != nil {
			if process.DiskInMB.IsSet {
				diskInMB = process.DiskInMB.Value
			}
			if process.MemoryInMB.IsSet {
				memoryInMB = process.MemoryInMB.Value
			}
		}
	}

	jobTimeout := args.Config["timeout"]
	var jt time.Duration = 30 * time.Minute

	if jobTimeout != "" {
		jt, err = time.ParseDuration(jobTimeout)
		if err != nil {
			return nil, fmt.Errorf("cftasks error: parsing job timeout")
		}
	}

	tickerTimeout := time.NewTicker(jt)

	noaaStreamer := NewNOAAStreamer(
		sess.NOAA(), sess.ConfigStore(),
		reportingWriter{buffer: output, cb: cb, isError: false},
		reportingWriter{buffer: output, cb: cb, isError: true},
	)
	defer noaaStreamer.Close()
	go func() {
		err := noaaStreamer.StreamLogsTask(appGUID, args.JobName)
		if err != nil {
			log.Printf("[Error] error on streaming log: %s\n", err.Error())
		}

	}()

	task, _, err := client.CreateApplicationTask(appGUID, ccv3.Task{
		Command:    command,
		DiskInMB:   diskInMB,
		MemoryInMB: memoryInMB,
		Name:       args.JobName,
	})
	if err != nil {
		return nil, err
	}

	ticker := time.NewTicker(time.Duration(config.PollingInterval) * time.Second)
	defer ticker.Stop()

	err = e.polltask(client, task, appGUID, ticker, tickerTimeout)
	resp := &types.ExecuteResponse{
		Output: output.Bytes(),
	}
	if err != nil {
		resp.Error = err.Error()
	}
	return resp, nil
}

func (e Executor) polltask(client *ccv3.Client, task ccv3.Task, appGuid string, ticker, tickerTimeout *time.Ticker) error {
	if task.State == constant.TaskSucceeded {
		return nil
	}
	if task.State == constant.TaskFailed {
		return fmt.Errorf("cftasks error: task has failed")
	}
	tasks, _, err := client.GetApplicationTasks(appGuid, ccv3.Query{
		Key:    ccv3.GUIDFilter,
		Values: []string{task.GUID},
	})
	if err != nil {
		return fmt.Errorf("cftasks error: when polling app: %s", err.Error())
	}
	if len(tasks) == 0 {
		return fmt.Errorf("cftasks error: when polling app the tasks has disappeared")
	}
	select {
	case <-ticker.C:
	case <-tickerTimeout.C:

		client.UpdateTaskCancel(tasks[0].GUID)
		return fmt.Errorf("cftasks error: timeout reached")

	}
	return e.polltask(client, tasks[0], appGuid, ticker, tickerTimeout)
}
