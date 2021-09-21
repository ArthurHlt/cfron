package models

import "github.com/orange-cloudfoundry/cfron/clients"

type JobExecution struct {
	Executions []clients.Execution `json:"executions"`
	ExecStatus string              `json:"exec_status"`
	Job        clients.Job         `json:"job"`
}
