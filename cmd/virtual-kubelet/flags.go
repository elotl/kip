package main

import "github.com/spf13/pflag"

type ServerConfig struct {
	DebugServer bool
}

func (c *ServerConfig) FlagSet() *pflag.FlagSet {
	flags := pflag.NewFlagSet("serverconfig", pflag.ContinueOnError)
	flags.BoolVar(&c.DebugServer, "debug-server", c.DebugServer, "Enable a listener in the server for inspecting internal milpa structures.")
	return flags
}
