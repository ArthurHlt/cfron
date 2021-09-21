package configs

import "github.com/kelseyhightower/envconfig"

type Config struct {
	ApiUrl            string `required:"true" split_words:"true"`
	User              string `default:"admin" split_words:"true"`
	Password          string `split_words:"true"`
	ClientId          string `split_words:"true"`
	ClientSecret      string `split_words:"true"`
	SkipSslValidation bool   `split_words:"true"`
	PollingInterval   uint   `default:"10" split_words:"true"`
}

var configStore *Config

func GetConfig() (*Config, error) {
	if configStore != nil {
		return configStore, nil
	}
	configStore = &Config{}
	err := envconfig.Process("cf", configStore)
	if err != nil {
		return nil, err
	}
	return configStore, nil
}
