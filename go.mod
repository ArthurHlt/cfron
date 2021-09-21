module github.com/orange-cloudfoundry/cfron

go 1.16

replace github.com/SermoDigital/jose => github.com/SermoDigital/jose v0.9.2-0.20161205224733-f6df55f235c2

replace github.com/mailru/easyjson => github.com/mailru/easyjson v0.0.0-20180323154445-8b799c424f57

replace github.com/cloudfoundry/sonde-go => github.com/cloudfoundry/sonde-go v0.0.0-20171206171820-b33733203bb4

replace code.cloudfoundry.org/go-log-cache => code.cloudfoundry.org/go-log-cache v1.0.1-0.20200316170138-f466e0302c34

require (
	code.cloudfoundry.org/bytefmt v0.0.0-20210608160410-67692ebc98de
	code.cloudfoundry.org/cli v7.1.0+incompatible
	code.cloudfoundry.org/lager v2.0.0+incompatible
	github.com/alecthomas/kingpin v2.2.6+incompatible
	github.com/alecthomas/units v0.0.0-20210208195552-ff826a37aa15 // indirect
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2
	github.com/cloudfoundry-community/go-cf-clients-helper v1.0.2
	github.com/cloudfoundry/noaa v2.1.0+incompatible
	github.com/cloudfoundry/sonde-go v0.0.0-20200416163440-a42463ba266b
	github.com/distribworks/dkron/v3 v3.1.7
	github.com/foolin/goview v0.3.0
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-plugin v1.4.2 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/nicklaw5/go-respond v0.0.0-20190722175312-54f5cd3d2246
	github.com/pivotal-cf/brokerapi/v8 v8.1.0
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.29.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/oauth2 v0.0.0-20210615190721-d04028783cf1
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.4.0
)
