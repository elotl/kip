package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/spf13/cobra"
)

func dump(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Not enough arguments")
		fatal("Usage: milpactl dump <all|stack|PodController|NodeController>")
	}
	controller := args[0]
	client, conn, err := getMilpaClient(cmd.InheritedFlags(), true)
	dieIfError(err, "Failed to create milpa client")
	defer conn.Close()

	dumpRequest := &clientapi.DumpRequest{
		Kind: []byte(controller),
	}
	reply, err := client.Dump(context.Background(), dumpRequest)
	dieIfError(err, "Could not dumpresource")
	dieIfReplyError("Dump", reply)
	fmt.Println(string(reply.Body))
}

func DumpCommand() *cobra.Command {

	var dumpCmd = &cobra.Command{
		Use:   "dump",
		Short: "Dump Controller Data",
		Long:  `Dump internal state of a controller`,
		Run: func(cmd *cobra.Command, args []string) {
			dump(cmd, args)
		},
	}

	return dumpCmd
}
