package metrics

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/mux"
	"github.com/orange-cloudfoundry/cfron/clients"
	"github.com/orange-cloudfoundry/cfron/errors"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

var gzipPool = sync.Pool{
	New: func() interface{} {
		return gzip.NewWriter(nil)
	},
}

type Metric struct {
	client *clients.APIClient
}

func NewMetric(client *clients.APIClient) *Metric {
	return &Metric{client: client}
}

func (m *Metric) GetTasksMetrics(rsp http.ResponseWriter, req *http.Request) {

	params := req.URL.Query()
	if len(params) == 0 {
		http.Error(rsp, "Missing params filter in query", http.StatusPreconditionFailed)
		return
	}

	name := ""
	metadata := make(map[string]string)

	for k, values := range params {
		if len(values) == 0 {
			continue
		}
		v := values[0]
		if v == "" {
			continue
		}
		if k == "name" {
			name = values[0]
			continue
		}
		metadata[k] = v
	}

	reqJobs := m.client.JobsApi.GetJobs(req.Context()).Sort("name")
	if len(metadata) > 0 {
		reqJobs = reqJobs.Metadata(metadata)
	}

	if name != "" {
		reqJobs = reqJobs.Q(name)
	}

	jobs, resp, errApi := reqJobs.Execute()
	err := errors.NewErrorFromClient(errApi)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if resp != nil {
			statusCode = resp.StatusCode
		}
		http.Error(rsp, fmt.Sprintf("Error from client: %v", err), statusCode)
		return
	}

	format := expfmt.Negotiate(req.Header)
	header := rsp.Header()
	header.Set("Content-Type", string(format))
	w := io.Writer(rsp)

	if gzipAccepted(req.Header) {
		header.Set("Content-Encoding", "gzip")
		gz := gzipPool.Get().(*gzip.Writer)
		defer gzipPool.Put(gz)

		gz.Reset(w)
		defer gz.Close()

		w = gz
	}
	counterSuccesses := make([]*dto.Metric, 0)
	counterErrores := make([]*dto.Metric, 0)
	gaugeCurrentes := make([]*dto.Metric, 0)
	for _, j := range jobs {
		labels := []*dto.LabelPair{
			{
				Name:  proto.String("name"),
				Value: proto.String(j.Name),
			},
			{
				Name:  proto.String("display_name"),
				Value: proto.String(*j.Displayname),
			},
		}
		for k, v := range *j.Metadata {
			labels = append(labels, &dto.LabelPair{
				Name:  proto.String(k),
				Value: proto.String(v),
			})
		}
		successCount := int32(0)
		if j.SuccessCount != nil {
			successCount = *j.SuccessCount
		}
		errorCount := int32(0)
		if j.ErrorCount != nil {
			errorCount = *j.ErrorCount
		}
		counterSuccesses = append(counterSuccesses, &dto.Metric{
			Label: labels,
			Counter: &dto.Counter{
				Value: proto.Float64(float64(successCount)),
			},
		})
		counterErrores = append(counterErrores, &dto.Metric{
			Label: labels,
			Counter: &dto.Counter{
				Value: proto.Float64(float64(errorCount)),
			},
		})
		status := float64(0)
		if j.Status != nil && *j.Status != "success" && *j.Status != "" {
			status = float64(1)
		}
		gaugeCurrentes = append(gaugeCurrentes, &dto.Metric{
			Label: labels,
			Gauge: &dto.Gauge{
				Value: proto.Float64(status),
			},
		})
	}

	counterType := dto.MetricType_COUNTER
	gaugeType := dto.MetricType_GAUGE

	enc := expfmt.NewEncoder(w, format)

	err = enc.Encode(&dto.MetricFamily{
		Name:   proto.String("cfron_task_success_total"),
		Help:   proto.String("Give total number if tasks in success until the beginning"),
		Type:   &counterType,
		Metric: counterSuccesses,
	})
	if err != nil && !strings.Contains(err.Error(), "broken pipe") {
		log.Warningf("Error when encoding exp fmt: %s", err.Error())
	}

	err = enc.Encode(&dto.MetricFamily{
		Name:   proto.String("cfron_task_error_total"),
		Help:   proto.String("Give total number if tasks in error until the beginning"),
		Type:   &counterType,
		Metric: counterErrores,
	})
	if err != nil && !strings.Contains(err.Error(), "broken pipe") {
		log.Warningf("Error when encoding exp fmt: %s", err.Error())
	}

	err = enc.Encode(&dto.MetricFamily{
		Name:   proto.String("cfron_current_status"),
		Help:   proto.String("Give current status of the task (0=success, 1=errror)"),
		Type:   &gaugeType,
		Metric: gaugeCurrentes,
	})
	if err != nil && !strings.Contains(err.Error(), "broken pipe") {
		log.Warningf("Error when encoding exp fmt: %s", err.Error())
	}

	if closer, ok := enc.(expfmt.Closer); ok {
		// This in particular takes care of the final "# EOF\n" line for OpenMetrics.
		closer.Close()
	}
}

func (m *Metric) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/tasks/metrics", m.GetTasksMetrics)
}

func gzipAccepted(header http.Header) bool {
	a := header.Get("Accept-Encoding")
	parts := strings.Split(a, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "gzip" || strings.HasPrefix(part, "gzip;") {
			return true
		}
	}
	return false
}
