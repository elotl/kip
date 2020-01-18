package cmd

import (
	"fmt"

	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

// This is here just for testing purposes
func setupIPForwarding(cmd *cobra.Command, args []string) {
	client, conn, err := getMilpaClient(cmd.InheritedFlags(), false)
	dieIfError(err, "Failed to create milpa client")
	defer conn.Close()

	req := &clientapi.SetupIPForwardingRequest{
		PodName: []byte(args[0]),
		Network: []byte("100.96.1.0/24"),
	}
	reply, err := client.SetupIPForwarding(context.Background(), req)
	dieIfError(err, "could not do setup ip forwarding")
	fmt.Printf("Milpa server reply: %d - %s\n", reply.Status, string(reply.Body))
}

func SetupIpForwardingCommand() *cobra.Command {
	var cobraCmd = &cobra.Command{
		Use:   "setup-ip-forwarding",
		Args:  cobra.ExactArgs(1),
		Short: "setup ip forarding",
		Long:  `setup cloud to forward traffic to a partcular pod`,
		Run: func(cmd *cobra.Command, args []string) {
			setupIPForwarding(cmd, args)
		},
	}
	return cobraCmd
}
