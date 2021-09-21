package models

import (
	"crypto/tls"
	"fmt"
	promodel "github.com/prometheus/common/model"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Listen         string          `yaml:"listen,omitempty"`
	EnableSSL      bool            `yaml:"enable_ssl,omitempty"`
	SSLCertificate tls.Certificate `yaml:"-"`
	TLSPEM         TLSPem          `yaml:"tls_pem,omitempty"`
	Log            Log             `yaml:"log"`
	Broker         Broker          `yaml:"broker"`
	Dkron          Dkron           `yaml:"dkron"`
	DevMode        bool            `yaml:"dev_mode"`
}

type Dkron struct {
	Endpoint string `yaml:"endpoint"`
}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain Config
	err := unmarshal((*plain)(c))
	if err != nil {
		return err
	}
	if c.EnableSSL {
		if c.TLSPEM.PrivateKey == "" || c.TLSPEM.CertChain == "" {
			return fmt.Errorf("Error parsing PEM blocks of router.tls_pem, missing cert or key.")
		}

		certificate, err := tls.X509KeyPair([]byte(c.TLSPEM.CertChain), []byte(c.TLSPEM.PrivateKey))
		if err != nil {
			errMsg := fmt.Sprintf("Error loading key pair: %s", err.Error())
			return fmt.Errorf(errMsg)
		}
		c.SSLCertificate = certificate
	}
	if c.Listen == "" {
		c.Listen = "0.0.0.0:9000"
	}

	return nil
}

type Log struct {
	Level   string `yaml:"level"`
	NoColor bool   `yaml:"no_color"`
	InJson  bool   `yaml:"in_json"`
}

func (c *Log) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain Log
	err := unmarshal((*plain)(c))
	if err != nil {
		return err
	}
	log.SetFormatter(&log.TextFormatter{
		DisableColors: c.NoColor,
	})
	if c.Level != "" {
		lvl, err := log.ParseLevel(c.Level)
		if err != nil {
			return err
		}
		log.SetLevel(lvl)
	}
	if c.InJson {
		log.SetFormatter(&log.JSONFormatter{})
	}

	return nil
}

type TLSPem struct {
	CertChain  string `yaml:"cert_chain"`
	PrivateKey string `yaml:"private_key"`
}

type Broker struct {
	ServiceName string `yaml:"service_name"`
	ServiceID   string `yaml:"service_id"`

	PlanName string `yaml:"plan_name"`
	PlanID   string `yaml:"plan_id"`

	Username string `yaml:"username"`
	Password string `yaml:"password"`

	OriginUrl   string            `yaml:"origin_url"`
	MinInterval promodel.Duration `yaml:"min_interval"`
}

func (c *Config) Initialize(configYAML []byte) error {
	return yaml.Unmarshal(configYAML, &c)
}

func InitConfigFromFile(filename string) (*Config, error) {
	c := &Config{}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = c.Initialize(b)
	if err != nil {
		return nil, err
	}

	return c, nil
}
