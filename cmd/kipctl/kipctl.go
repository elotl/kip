/*
Copyright 2020 Elotl Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	goflag "flag"
	"fmt"
	"log"
	"os"

	"github.com/elotl/kip/cmd/kipctl/cmd"
	"github.com/elotl/kip/pkg/util"
	"github.com/spf13/cobra"
)

const (
	cliName         = "kipctl"
	cliDescription  = "Command line client for Kip provider"
	defaultEndpoint = "localhost:54555"
)

var (
	rootCmd = &cobra.Command{
		Use:        cliName,
		Short:      cliDescription,
		SuggestFor: []string{"kipctl"},
		Run: func(cmd *cobra.Command, args []string) {
			_ = goflag.CommandLine.Parse([]string{})
			if version {
				fmt.Printf("%s version %s\n", cliName, util.Version())
				os.Exit(0)
			}
		},
	}
	endpoints []string // used by rootCmd
	version   bool
)

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().AddGoFlagSet(goflag.CommandLine)
	cmd.SetupKnownTypes()
}

func main() {
	rootCmd.PersistentFlags().StringSliceVar(&endpoints, "endpoints", []string{defaultEndpoint}, "comma separated list of server IP and Port ('ip:port') endpoints to connect to")

	rootCmd.AddCommand(cmd.AttachCommand())
	rootCmd.AddCommand(cmd.CreateCommand())
	rootCmd.AddCommand(cmd.DeleteCommand())
	rootCmd.AddCommand(cmd.DeployCommand())
	rootCmd.AddCommand(cmd.DumpCommand())
	rootCmd.AddCommand(cmd.ExecCommand())
	rootCmd.AddCommand(cmd.GetCommand())
	rootCmd.AddCommand(cmd.LogsCommand())
	rootCmd.AddCommand(cmd.UpdateCommand())
	rootCmd.AddCommand(cmd.VersionCommand())

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to execute command: %v", err)
	}
}
