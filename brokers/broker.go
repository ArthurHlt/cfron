package brokers

import (
	"code.cloudfoundry.org/lager"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/distribworks/dkron/v3/extcron"
	"github.com/orange-cloudfoundry/cfron/clients"
	"github.com/orange-cloudfoundry/cfron/models"
	"github.com/pivotal-cf/brokerapi/v8"
	"github.com/pivotal-cf/brokerapi/v8/domain"
	"github.com/pivotal-cf/brokerapi/v8/middlewares"
	"github.com/robfig/cron/v3"
	"net/http"
	"os"
	"strings"
)

type DkronBroker struct {
	brokerConfig    models.Broker
	client          *clients.APIClient
	cronParser      cron.ScheduleParser
	executorFactory func(appGuid string, bindParams models.BindParams) (string, map[string]string)
}

func NewDkronBroker(brokerConfig models.Broker, client *clients.APIClient) *DkronBroker {
	return &DkronBroker{
		brokerConfig:    brokerConfig,
		client:          client,
		executorFactory: CFTasksExecutorFactory,
		cronParser:      extcron.NewParser(),
	}
}

func (d *DkronBroker) SetExecutorFactory(executorFactory func(appGuid string, bindParams models.BindParams) (string, map[string]string)) {
	d.executorFactory = executorFactory
}

func (d *DkronBroker) Handler() http.Handler {
	lag := lager.NewLogger("broker")
	lag.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
	return brokerapi.New(d, lag, brokerapi.BrokerCredentials{
		Username: d.brokerConfig.Username,
		Password: d.brokerConfig.Password,
	})
}

func (d DkronBroker) Services(ctx context.Context) ([]domain.Service, error) {
	return []domain.Service{
		{
			ID:                   d.brokerConfig.ServiceID,
			Name:                 d.brokerConfig.ServiceName,
			Description:          "A service broker to create cron task",
			Bindable:             true,
			InstancesRetrievable: true,
			BindingsRetrievable:  true,
			Tags:                 []string{"cron", "periodic"},
			PlanUpdatable:        false,
			Plans: []domain.ServicePlan{
				{
					ID:          d.brokerConfig.PlanID,
					Name:        d.brokerConfig.PlanName,
					Description: "Plan to create a cron",
					Free:        domain.FreeValue(true),
					Bindable:    domain.BindableValue(true),
					Metadata: &domain.ServicePlanMetadata{
						DisplayName: d.brokerConfig.PlanName,
					},
					Schemas:                nil,
					PlanUpdatable:          nil,
					MaximumPollingDuration: nil,
					MaintenanceInfo:        nil,
				},
			},
			Requires: []domain.RequiredPermission{},
			Metadata: &domain.ServiceMetadata{
				DisplayName:         d.brokerConfig.ServiceName,
				ImageUrl:            "",
				LongDescription:     "A service broker to create cron task over dkron",
				ProviderDisplayName: "Orange",
				DocumentationUrl:    "",
				SupportUrl:          "https://github.com/orange-cloudfoundry/cfron",
			},
			DashboardClient:     nil,
			AllowContextUpdates: true,
		},
	}, nil
}

func ctxToUser(ctx context.Context) (models.CFPlatform, error) {
	resp := ctx.Value(middlewares.OriginatingIdentityKey)
	if resp == nil {
		return models.CFPlatform{
			UserId: "unknown",
		}, nil
	}

	platformRaw := resp.(string)
	platformSplit := strings.SplitN(platformRaw, " ", 2)
	if len(platformSplit) != 2 {
		return models.CFPlatform{
			UserId: "unknown",
		}, nil
	}

	rawUser := platformSplit[1]
	b, err := base64.StdEncoding.DecodeString(rawUser)
	if err != nil {
		return models.CFPlatform{}, err
	}
	var platform models.CFPlatform
	err = json.Unmarshal(b, &platform)
	if err != nil {
		return models.CFPlatform{}, err
	}
	return platform, nil
}

func StringPtr(s string) *string {
	return &s
}

func (d DkronBroker) makeDashboardUrl(instanceID string) string {
	return fmt.Sprintf("%s/dashboard/tasks?instance_id=%s", d.brokerConfig.OriginUrl, instanceID)
}

func (d DkronBroker) makeJobUrl(bindingID, instanceID string) string {
	return fmt.Sprintf("%s/dashboard/tasks?instance_id=%s&name=%s", d.brokerConfig.OriginUrl, instanceID, bindingID)
}

func CFTasksExecutorFactory(appGuid string, bindParams models.BindParams) (string, map[string]string) {
	return "cftasks", map[string]string{
		"command":  bindParams.Command,
		"memory":   bindParams.Memory,
		"disk":     bindParams.Disk,
		"timeout":  bindParams.Timeout,
		"app_guid": appGuid,
	}
}

func ShellExecutorFactory(appGuid string, bindParams models.BindParams) (string, map[string]string) {
	return "shell", map[string]string{
		"shell":   "true",
		"command": bindParams.Command,
		"timeout": bindParams.Timeout,
	}
}
