package brokers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/orange-cloudfoundry/cfron/errors"
	"github.com/orange-cloudfoundry/cfron/models"
	"github.com/pivotal-cf/brokerapi/v8/domain"
)

func (d DkronBroker) Provision(ctx context.Context, instanceID string, details domain.ProvisionDetails, asyncAllowed bool) (domain.ProvisionedServiceSpec, error) {
	return domain.ProvisionedServiceSpec{
		IsAsync:       false,
		AlreadyExists: false,
		DashboardURL:  d.makeDashboardUrl(instanceID),
		OperationData: "",
		Metadata:      domain.InstanceMetadata{},
	}, nil
}

func (d DkronBroker) Deprovision(ctx context.Context, instanceID string, details domain.DeprovisionDetails, asyncAllowed bool) (domain.DeprovisionServiceSpec, error) {
	return domain.DeprovisionServiceSpec{
		IsAsync:       false,
		OperationData: "",
	}, nil
}

func (d DkronBroker) GetInstance(ctx context.Context, instanceID string, details domain.FetchInstanceDetails) (domain.GetInstanceDetailsSpec, error) {
	return domain.GetInstanceDetailsSpec{
		DashboardURL: d.makeDashboardUrl(instanceID),
		Metadata:     domain.InstanceMetadata{},
	}, nil
}

func (d DkronBroker) Update(ctx context.Context, instanceID string, details domain.UpdateDetails, asyncAllowed bool) (domain.UpdateServiceSpec, error) {

	var cfContext models.CFContext
	err := json.Unmarshal(details.RawContext, &cfContext)
	if err != nil && len(details.RawContext) > 0 {
		return domain.UpdateServiceSpec{}, fmt.Errorf("Error when loading context: %s", err.Error())
	}

	jobs, _, errApi := d.client.JobsApi.GetJobs(ctx).Metadata(map[string]string{
		"instance_id": instanceID,
	}).Execute()

	err = errors.NewErrorFromClient(errApi)
	if err != nil {
		return domain.UpdateServiceSpec{}, err
	}

	for _, job := range jobs {
		var metadata map[string]string
		if job.Metadata == nil {
			metadata = make(map[string]string)
			job.Metadata = &metadata
		} else {
			metadata = *job.Metadata
		}

		metadata["space_name"] = cfContext.SpaceName
		metadata["space_guid"] = cfContext.SpaceGUID
		metadata["organization_name"] = cfContext.OrganizationName
		metadata["organization_guid"] = cfContext.OrganizationGUID
		metadata["platform"] = cfContext.Platform
		metadata["instance_name"] = cfContext.InstanceName
		_, _, errApi := d.client.JobsApi.CreateOrUpdateJob(ctx).Body(job).Execute()
		err = errors.NewErrorFromClient(errApi)
		if err != nil {
			return domain.UpdateServiceSpec{}, err
		}
	}

	return domain.UpdateServiceSpec{
		IsAsync:       false,
		DashboardURL:  d.makeDashboardUrl(instanceID),
		OperationData: "",
		Metadata:      domain.InstanceMetadata{},
	}, nil
}

func (d DkronBroker) LastOperation(ctx context.Context, instanceID string, details domain.PollDetails) (domain.LastOperation, error) {
	return domain.LastOperation{
		State:       domain.Succeeded,
		Description: "",
	}, nil
}
