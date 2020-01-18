package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/wsstream"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	logsLines      int
	logsLimitBytes int
	logsFollow     bool
	logsUnitName   string
)

func getLogs(cmd *cobra.Command, args []string) {
	resourceName := args[0]
	if logsLimitBytes == 0 && logsLines == 0 {
		logsLines = 20
	}

	client, conn, err := getMilpaClient(cmd.InheritedFlags(), false)
	dieIfError(err, "Failed to create milpa client")
	defer conn.Close()

	if logsFollow {
		tailLogs(client, resourceName, logsUnitName)
	} else {
		logsRequest := &clientapi.LogsRequest{
			ResourceName: resourceName,
			ItemName:     logsUnitName,
			Lines:        int32(logsLines),
			Limitbytes:   int32(logsLimitBytes)}
		reply, err := client.GetLogs(context.Background(), logsRequest)
		dieIfError(err, "Could not get %s logs", resourceName)
		dieIfReplyError("Logs", reply)
		obj, err := api.Decode(reply.Body)
		dieIfError(err, "")

		logfile, ok := obj.(*api.LogFile)
		if !ok {
			fatal(
				"Got back unknown object type.\nObject:\n%#v",
				string(reply.Body),
			)
		}
		fmt.Printf("Resource %s (item %s) logs:\n%s",
			resourceName, logsUnitName, string(logfile.Content))
	}
}

func tailLogs(client clientapi.MilpaClient, resourceName, unitName string) {
	slr := &clientapi.StreamLogsRequest{
		Pod:      resourceName,
		Unit:     unitName,
		Metadata: false,
	}
	stream, err := client.StreamLogs(context.Background(), slr)
	dieIfError(err, "Error streaming logs")
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error streaming logs", err)
			break
		}
		if len(msg.Data) > 0 {
			unpackStreamMsg(msg.Data)
		}
	}
}

func unpackStreamMsg(f []byte) {
	c, m, err := wsstream.UnpackMessage(f)
	if err != nil {
		fmt.Println("Corrupted message", err)
		return
	}
	if c == wsstream.StderrChan {
		fmt.Fprint(os.Stderr, string(m))
	} else {
		fmt.Fprint(os.Stdout, string(m))
	}
}

func LogsCommand() *cobra.Command {
	var logsCmd = &cobra.Command{
		Use:   "logs resource_name [-u unit]",
		Short: "Get logs",
		Long: `Get logs of a given unit in a pod or get milpa agent logs from a node.
Milpa will save the tail logs of deleted resources and allow them to be queried for up to 1 hour.`,
		Example: "Pod Logs: milpactl logs mypod -u unitname --lines 25\nNode Logs: milpactl logs node-uuid",
		Args:    cobra.ExactArgs(1),
		Run:     getLogs,
	}
	logsCmd.Flags().StringVarP(&logsUnitName, "unit", "u", "", "Unit name. If empty the first unit in the pod will be used")
	logsCmd.Flags().BoolVarP(&logsFollow, "follow", "f", false, "Follow logs")
	logsCmd.Flags().IntVar(&logsLines, "lines", 0, "Number of lines to retrieve")
	logsCmd.Flags().IntVar(&logsLimitBytes, "limit-bytes", 0, "Limit length of logs")
	return logsCmd
}
