package cftasks

import (
	"code.cloudfoundry.org/cli/util/configv3"
	"fmt"
	noaaconsumer "github.com/cloudfoundry/noaa/consumer"
	"github.com/cloudfoundry/sonde-go/events"
	"io"
	"strings"
	"time"
)

const LogTimestampFormat = "2006-01-02T15:04:05.00-0700"

type NOAAStreamer struct {
	consumer     *noaaconsumer.Consumer
	configStore  *configv3.Config
	writerStdout io.Writer
	writerStderr io.Writer
}

func NewNOAAStreamer(consumer *noaaconsumer.Consumer, configStore *configv3.Config, writerStdout, writerStderr io.Writer) *NOAAStreamer {
	return &NOAAStreamer{
		consumer:     consumer,
		configStore:  configStore,
		writerStdout: writerStdout,
		writerStderr: writerStderr,
	}
}

func (c NOAAStreamer) Close() error {
	return c.consumer.Close()
}

func (c NOAAStreamer) StreamLogsTask(appGUID string, jobName string) error {

	envChan, errChan := c.consumer.StreamWithoutReconnect(appGUID, c.configStore.AccessToken())

	jobSourceType := fmt.Sprintf("APP/TASK/%s", jobName)
	for {
		select {
		case env := <-envChan:
			if env.GetLogMessage() == nil {
				continue
			}
			logMsg := env.GetLogMessage()
			if logMsg.GetSourceType() != jobSourceType {
				continue
			}
			t := time.Unix(0, logMsg.GetTimestamp()).In(time.Local).Format(LogTimestampFormat)

			header := fmt.Sprintf("%s ",
				t,
			)
			message := string(logMsg.GetMessage())
			writer := c.writerStdout
			if logMsg.GetMessageType() != events.LogMessage_OUT {
				writer = c.writerStderr
				header += "ERR "
			} else {
				header += "OUT "
			}
			for _, line := range strings.Split(message, "\n") {
				writer.Write([]byte(fmt.Sprintf("%s%s\n", header, strings.TrimRight(line, "\r\n"))))
			}

		case err := <-errChan:
			return err
		}
	}

	return nil
}
