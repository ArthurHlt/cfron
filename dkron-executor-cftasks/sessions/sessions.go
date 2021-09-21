package sessions

import (
	"os"

	clients "github.com/cloudfoundry-community/go-cf-clients-helper"
	"github.com/orange-cloudfoundry/cfron/dkron-executor-cftasks/configs"
)

func getEnvOrDefault(key, defaultVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	return val
}

func GetSession() (*clients.Session, error) {
	configExecutor, err := configs.GetConfig()
	if err != nil {
		return nil, err
	}
	config := clients.Config{
		Endpoint:          configExecutor.ApiUrl,
		User:              configExecutor.User,
		Password:          configExecutor.Password,
		CFClientID:        configExecutor.ClientId,
		CFClientSecret:    configExecutor.ClientSecret,
		SkipSslValidation: configExecutor.SkipSslValidation,
	}
	return clients.NewSession(config)
}
