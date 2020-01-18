package main

import (
	goflag "flag"
	"fmt"
	"log"
	"os"

	"github.com/elotl/cloud-instance-provider/cmd/milpactl/cmd"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/spf13/cobra"
)

const (
	cliName         = "milpactl"
	cliDescription  = "Command line client for Milpa framework."
	defaultEndpoint = "localhost:54555"
)

var (
	rootCmd = &cobra.Command{
		Use:        cliName,
		Short:      cliDescription,
		SuggestFor: []string{"milpactl"},
		Run: func(cmd *cobra.Command, args []string) {
			goflag.CommandLine.Parse([]string{})
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
