package main

import (
	dkplugin "github.com/distribworks/dkron/v3/plugin"
	"github.com/hashicorp/go-plugin"
	"github.com/orange-cloudfoundry/cfron/dkron-executor-cftasks/cftasks"
)

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: dkplugin.Handshake,
		Plugins: map[string]plugin.Plugin{
			"executor": &dkplugin.ExecutorPlugin{Executor: &cftasks.Executor{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
