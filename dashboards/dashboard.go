package dashboards

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/foolin/goview"
	"github.com/gorilla/mux"
	"github.com/nicklaw5/go-respond"
	"github.com/orange-cloudfoundry/cfron/clients"
	"github.com/orange-cloudfoundry/cfron/errors"
	"github.com/orange-cloudfoundry/cfron/front"
	"github.com/orange-cloudfoundry/cfron/models"
)

type Dashboard struct {
	client *clients.APIClient
	gv     *goview.ViewEngine
}

func NewDashboard(client *clients.APIClient) *Dashboard {
	gv := goview.New(goview.Config{
		Root:         "templates",
		Extension:    ".gohtml",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        template.FuncMap{},
		DisableCache: true,
		Delims:       goview.Delims{Left: "{{", Right: "}}"},
	})
	gv.SetFileHandler(func(config goview.Config, tplFile string) (content string, err error) {
		b, err := front.Templates.ReadFile(filepath.Join(config.Root, tplFile) + config.Extension)
		if err != nil {
			return "", err
		}
		return string(b), nil
	})
	return &Dashboard{
		client: client,
		gv:     gv,
	}
}

func (d *Dashboard) Index(w http.ResponseWriter, req *http.Request) {

	err := d.gv.Render(w, http.StatusOK, "index", goview.M{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Render index error: %v!", err), http.StatusInternalServerError)
		return
	}
}

func (d *Dashboard) Tasks(w http.ResponseWriter, req *http.Request) {
	name := ""
	metadata := make(map[string]string)
	for key, values := range req.URL.Query() {
		if len(values) == 0 || values[0] == "" {
			continue
		}
		if key == "name" {
			name = values[0]
			continue
		}
		metadata[key] = values[0]
	}
	err := d.gv.Render(w, http.StatusOK, "tasks", goview.M{
		"metadata": metadata,
		"name":     name,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Render index error: %v!", err), http.StatusInternalServerError)
		return
	}
}

func (d *Dashboard) GetTasks(w http.ResponseWriter, req *http.Request) {
	name := ""
	metadata := make(map[string]string)
	for key, values := range req.URL.Query() {
		if len(values) == 0 || values[0] == "" {
			continue
		}
		if key == "name" {
			name = values[0]
			continue
		}
		metadata[key] = values[0]

	}
	reqJobs := d.client.JobsApi.GetJobs(req.Context()).Sort("name")
	if len(metadata) > 0 {
		reqJobs = reqJobs.Metadata(metadata)
	}

	if name != "" {
		reqJobs = reqJobs.Q(name)
	}
	if len(metadata) == 0 && name == "" {
		// send empty list if no criteria
		respond.NewResponse(w).Ok([]string{})
		return
	}
	jobs, resp, errApi := reqJobs.Execute()
	err := errors.NewErrorFromClient(errApi)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if resp != nil {
			statusCode = resp.StatusCode
		}
		http.Error(w, fmt.Sprintf("Error from client: %v", err), statusCode)
		return
	}

	jobsExec := make([]models.JobExecution, len(jobs))
	for i, job := range jobs {
		executions, err := d.GetTaskExecutionsByName(req.Context(), job.Name)
		if err != nil {
			if errStatus, ok := err.(*errors.ErrWithStatusCode); ok {
				http.Error(w, fmt.Sprintf("Error from client: %v", errStatus.Msg), errStatus.StatusCode)
				return
			}
			http.Error(w, fmt.Sprintf("Error from client: %v", err), http.StatusInternalServerError)
			return
		}
		execStatus := "finished"
		if len(executions) == 0 {
			execStatus = "not_started"
		}
		if len(executions) > 0 && executions[0].FinishedAt.IsZero() {
			execStatus = "running"
		}

		jobsExec[i] = models.JobExecution{
			Executions: executions,
			ExecStatus: execStatus,
			Job:        job,
		}
	}
	respond.NewResponse(w).Ok(jobsExec)
}

func (d *Dashboard) GetTaskExecutionsByName(ctx context.Context, jobName string) ([]clients.Execution, error) {
	executions, resp, errApi := d.client.ExecutionsApi.ListExecutionsByJob(ctx, jobName).Execute()
	err := errors.NewErrorFromClient(errApi)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if resp != nil {
			statusCode = resp.StatusCode
		}
		return []clients.Execution{}, errors.NewErrorWithStatusCode(err.Error(), statusCode)
	}

	return executions, nil
}

func (d *Dashboard) GetTaskExecutions(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobName := vars["job_name"]
	execState, err := d.GetTaskExecutionsByName(req.Context(), jobName)
	if err != nil {
		if errStatus, ok := err.(*errors.ErrWithStatusCode); ok {
			http.Error(w, fmt.Sprintf("Error from client: %v", errStatus.Msg), errStatus.StatusCode)
			return
		}
		http.Error(w, fmt.Sprintf("Error from client: %v", err), http.StatusInternalServerError)
		return
	}

	respond.NewResponse(w).Ok(execState)
}

func (d *Dashboard) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/static/").Handler(http.FileServer(http.FS(front.Static)))
	router.HandleFunc("/", d.Index)
	router.HandleFunc("/dashboard", d.Index)
	router.HandleFunc("/dashboard/api/v1/tasks/{job_name}/executions", d.GetTaskExecutions)
	router.HandleFunc("/dashboard/api/v1/tasks", d.GetTasks)
	router.HandleFunc("/dashboard/tasks", d.Tasks)
}
