package main

import "github.com/spf13/pflag"

type ServerConfig struct {
	DebugServer        bool
	NetworkAgentSecret string
}

func (c *ServerConfig) FlagSet() *pflag.FlagSet {
	flags := pflag.NewFlagSet("serverconfig", pflag.ContinueOnError)
	flags.BoolVar(&c.DebugServer, "debug-server", c.DebugServer, "Enable a listener in the server for inspecting internal milpa structures.")
	flags.StringVar(&c.NetworkAgentSecret, "network-agent-secret", c.NetworkAgentSecret, "Service account secret for the cell network agent, in the form of <namespace>/<name>")
	return flags
}
