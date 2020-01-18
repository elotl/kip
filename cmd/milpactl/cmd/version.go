package cmd

import (
	"fmt"

	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

func getServerVersion(cmd *cobra.Command, args []string) {
	client, conn, err := getMilpaClient(cmd.InheritedFlags(), false)
	dieIfError(err, "Failed to create milpa client")
	defer conn.Close()
	versionRequest := &clientapi.VersionRequest{}
	reply, err := client.GetVersion(context.Background(), versionRequest)
	dieIfError(err, "could not gather grpc server version")
	fmt.Printf("Milpa server version: %s\n", string(reply.VersionInfo))
}

func VersionCommand() *cobra.Command {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Get version",
		Long:  `Get framework version`,
		Run: func(cmd *cobra.Command, args []string) {
			getServerVersion(cmd, args)
		},
	}
	return versionCmd
}
