package brokers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/distribworks/dkron/v3/extcron"
	"net/http"
	"time"

	"github.com/orange-cloudfoundry/cfron/clients"
	"github.com/orange-cloudfoundry/cfron/errors"
	"github.com/orange-cloudfoundry/cfron/models"
	"github.com/pivotal-cf/brokerapi/v8/domain"
	"github.com/pivotal-cf/brokerapi/v8/domain/apiresponses"
)

func (d DkronBroker) Bind(ctx context.Context, instanceID, bindingID string, details domain.BindDetails, asyncAllowed bool) (domain.Binding, error) {

	cfPlatform, err := ctxToUser(ctx)
	if err != nil {
		return domain.Binding{}, fmt.Errorf("Error when loading platform: %s", err.Error())
	}

	appGuid := details.AppGUID
	if details.BindResource != nil {
		appGuid = details.BindResource.AppGuid
	}
	var cfContext models.CFContext
	err = json.Unmarshal(details.RawContext, &cfContext)
	if err != nil && len(details.RawContext) > 0 {
		return domain.Binding{}, fmt.Errorf("Error when loading context: %s", err.Error())
	}

	var bindParams models.BindParams
	err = json.Unmarshal(details.RawParameters, &bindParams)
	if err != nil && len(details.RawParameters) > 0 {
		return domain.Binding{}, fmt.Errorf("Error when loading params: %s", err.Error())
	}

	if bindParams.Timeout == "" {
		bindParams.Timeout = "30m"
	}

	if bindParams.Schedule == "" {
		return domain.Binding{}, errors.NewErrorResponse("schedule param is mandatory", http.StatusUnprocessableEntity, "bind-params")
	}
	if bindParams.Command == "" {
		return domain.Binding{}, errors.NewErrorResponse("command param is mandatory", http.StatusUnprocessableEntity, "bind-params")
	}

	err = d.checkSchedule(bindParams.Schedule)
	if err != nil {
		return domain.Binding{}, errors.NewErrorResponse("schedule param is wrong: "+err.Error(), http.StatusUnprocessableEntity, "check-schedule")
	}

	displayName := bindParams.Name
	if displayName == "" {
		displayName = bindingID
	}
	metadata := map[string]string{
		"instance_name":     cfContext.InstanceName,
		"instance_id":       instanceID,
		"app_guid":          appGuid,
		"space_name":        cfContext.SpaceName,
		"space_guid":        cfContext.SpaceGUID,
		"organization_name": cfContext.OrganizationName,
		"organization_guid": cfContext.OrganizationGUID,
		"platform":          cfContext.Platform,
	}
	owner := cfPlatform.UserId

	executor, executorConfig := d.executorFactory(appGuid, bindParams)

	concurrency := "forbid"
	if bindParams.AllowConcurrency {
		concurrency = "allow"
	}
	_, _, errApi := d.client.JobsApi.CreateOrUpdateJob(ctx).Body(clients.Job{
		Name:           bindingID,
		Displayname:    &displayName,
		Schedule:       bindParams.Schedule,
		Owner:          &owner,
		Retries:        bindParams.Retries,
		Executor:       StringPtr(executor),
		ExecutorConfig: &executorConfig,
		Concurrency:    &concurrency,
		Metadata:       &metadata,
		Status:         nil,
	}).Execute()

	err = errors.NewErrorFromClient(errApi)
	if err != nil {
		return domain.Binding{}, err
	}

	return domain.Binding{
		Credentials: models.BindCredentials{
			Name:           bindingID,
			DisplayName:    displayName,
			Schedule:       bindParams.Schedule,
			Executor:       "cftasks",
			ExecutorConfig: executorConfig,
			Metadata:       metadata,
			Retries:        bindParams.Retries,
			Status:         "",
			Owner:          owner,
			DashboardUrl:   d.makeJobUrl(bindingID, instanceID),
		},
	}, nil
}

func (d DkronBroker) Unbind(ctx context.Context, instanceID, bindingID string, details domain.UnbindDetails, asyncAllowed bool) (domain.UnbindSpec, error) {

	_, resp, errApi := d.client.JobsApi.DeleteJob(ctx, bindingID).Execute()
	err := errors.NewErrorFromClient(errApi)
	if err != nil && resp.StatusCode == http.StatusNotFound {
		return domain.UnbindSpec{
			IsAsync:       false,
			OperationData: "",
		}, nil
	}
	if err != nil {
		return domain.UnbindSpec{
			IsAsync:       false,
			OperationData: "",
		}, err
	}

	return domain.UnbindSpec{
		IsAsync:       false,
		OperationData: "",
	}, nil
}

func (d DkronBroker) GetBinding(ctx context.Context, instanceID, bindingID string, details domain.FetchBindingDetails) (domain.GetBindingSpec, error) {
	job, resp, errApi := d.client.JobsApi.ShowJobByName(ctx, bindingID).Execute()
	err := errors.NewErrorFromClient(errApi)
	if resp.StatusCode == http.StatusNotFound {
		return domain.GetBindingSpec{}, apiresponses.ErrBindingNotFound
	}
	if err != nil {
		return domain.GetBindingSpec{}, errApi
	}
	displayName := ""
	if job.Displayname != nil {
		displayName = *job.Displayname
	}

	executorConfig := map[string]string{}
	if job.ExecutorConfig != nil {
		executorConfig = *job.ExecutorConfig
	}

	metadata := map[string]string{}
	if job.Metadata != nil {
		metadata = *job.Metadata
	}

	status := ""
	if job.Status != nil {
		status = *job.Status
	}

	owner := ""
	if job.Owner != nil {
		owner = *job.Owner
	}

	concurrency := "forbid"
	if job.Concurrency != nil {
		concurrency = *job.Concurrency
	}

	return domain.GetBindingSpec{
		Credentials: models.BindCredentials{
			Name:           bindingID,
			DisplayName:    displayName,
			Schedule:       job.Schedule,
			Executor:       "cftasks",
			ExecutorConfig: executorConfig,
			Metadata:       metadata,
			Retries:        job.Retries,
			Status:         status,
			Owner:          owner,
			Concurrency:    concurrency,
		},
	}, nil
}

func (d DkronBroker) LastBindingOperation(ctx context.Context, instanceID, bindingID string, details domain.PollDetails) (domain.LastOperation, error) {
	return domain.LastOperation{
		State:       domain.Succeeded,
		Description: "",
	}, nil
}

func (d DkronBroker) checkSchedule(scheduleStr string) error {
	timeNow := time.Now()
	schedule, err := d.cronParser.Parse(scheduleStr)
	if err != nil {
		return fmt.Errorf("Parse error: %v", err)
	}
	if _, ok := schedule.(extcron.SimpleSchedule); ok {
		return nil
	}

	interval := schedule.Next(timeNow).Sub(timeNow)

	if interval < time.Minute {
		interval = interval.Round(time.Second)
	}
	if interval > time.Minute {
		interval = interval.Round(time.Minute)
	}
	if interval < time.Duration(d.brokerConfig.MinInterval) {
		return fmt.Errorf("Minimum interval is %s and your interval is %s", d.brokerConfig.MinInterval, interval.String())
	}
	return nil
}
