package models

type BindParams struct {
	Name             string `json:"name"`
	Command          string `json:"command"`
	Schedule         string `json:"schedule"`
	Disk             string `json:"disk"`
	Memory           string `json:"memory"`
	Timeout          string `json:"timeout"`
	Retries          *int32 `json:"retries"`
	AllowConcurrency bool   `json:"allow_concurrency"`
}

type BindCredentials struct {
	Name           string            `json:"job_name"`
	DisplayName    string            `json:"display_name"`
	Schedule       string            `json:"schedule"`
	Executor       string            `json:"executor"`
	ExecutorConfig map[string]string `json:"executor_config"`
	Metadata       map[string]string `json:"metadata"`
	Retries        *int32            `json:"retries,omitempty"`
	Status         string            `json:"status,omitempty"`
	Owner          string            `json:"owner"`
	Concurrency    string            `json:"concurrency"`
	DashboardUrl   string            `json:"dashboard_url"`
}
