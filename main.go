package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/gorilla/mux"
	"github.com/prometheus/common/version"
	log "github.com/sirupsen/logrus"

	"github.com/orange-cloudfoundry/cfron/brokers"
	"github.com/orange-cloudfoundry/cfron/clients"
	"github.com/orange-cloudfoundry/cfron/dashboards"
	"github.com/orange-cloudfoundry/cfron/metrics"
	"github.com/orange-cloudfoundry/cfron/models"
)

var (
	configFile = kingpin.Flag("config", "Configuration File").Short('c').Default("config.yml").String()
)

func main() {
	kingpin.Version(version.Print("dkron-broker"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	c, err := models.InitConfigFromFile(*configFile)
	if err != nil {
		log.Fatal("Error loading config: ", err.Error())
	}

	apiClient, err := makeApiClient(c)
	if err != nil {
		log.Fatal("Error loading client for dkron: ", err.Error())
	}

	broker := brokers.NewDkronBroker(c.Broker, apiClient)
	if c.DevMode {
		broker.SetExecutorFactory(brokers.ShellExecutorFactory)
	}

	dahsboard := dashboards.NewDashboard(apiClient)
	metric := metrics.NewMetric(apiClient)

	rtr := mux.NewRouter()
	rtr.NewRoute().MatcherFunc(func(req *http.Request, m *mux.RouteMatch) bool {
		return strings.HasPrefix(req.URL.Path, "/broker/v2")
	}).Handler(http.StripPrefix("/broker", broker.Handler()))
	dahsboard.RegisterRoutes(rtr)
	metric.RegisterRoutes(rtr)

	srvSignal := make(chan os.Signal, 1)
	signal.Notify(srvSignal, syscall.SIGTERM, syscall.SIGINT)

	srvCtx, cancel := context.WithCancel(context.Background())

	go func() {
		<-srvSignal
		cancel()
	}()

	listener, err := makeListener(c)
	if err != nil {
		log.Fatal(err.Error())
	}
	srv := &http.Server{Handler: rtr}

	go func() {
		if err = srv.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %+s\n", err)
		}
	}()
	defer srv.Close()

	<-srvCtx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer func() {
		cancel()
	}()

	err = srv.Shutdown(ctxShutDown)
	if err != nil {
		log.Fatalf("server shutdown gracefully Failed: %s\n", err.Error())
	}
	log.Info("server gracefully shutdown")
}

func makeApiClient(c *models.Config) (*clients.APIClient, error) {
	clientConfig := clients.NewConfiguration()
	u, err := url.Parse(c.Dkron.Endpoint)
	if err != nil {
		return nil, err
	}
	clientConfig.Host = u.Host
	clientConfig.Scheme = u.Scheme
	return clients.NewAPIClient(clientConfig), nil
}

func makeListener(c *models.Config) (net.Listener, error) {
	listenAddr := c.Listen
	if !c.EnableSSL {
		log.Infof("Listen %s without tls ...", listenAddr)
		return net.Listen("tcp", listenAddr)
	}
	log.Infof("Listen %s with tls ...", listenAddr)
	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		rootCAs = nil
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{c.SSLCertificate},
		ClientCAs:    rootCAs,
	}

	tlsConfig.BuildNameToCertificate()
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, err
	}
	return tls.NewListener(listener, tlsConfig), nil
}
