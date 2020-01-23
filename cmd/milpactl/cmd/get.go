package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/milpactl"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

func get(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		msg := "Not enough arguments\nUsage: milpactl get (pod|node|service|resource) [Name]"
		fatal(msg)
	}
	kind := milpactl.CleanupResourceName(args[0])
	var name string
	if len(args) > 1 {
		name = args[1]
	}
	if !util.StringInSlice(kind, getTypes) {
		fatal("Illegal resource type for GET: %s", kind)
	}

	client, conn, err := getMilpaClient(cmd.InheritedFlags(), false)
	dieIfError(err, "Failed to create milpa client")
	defer conn.Close()

	getRequest := &clientapi.GetRequest{
		Kind: []byte(kind),
		Name: []byte(name),
	}
	reply, err := client.Get(context.Background(), getRequest)
	dieIfError(err, "Could not get resource")
	dieIfReplyError("Get", reply)
	printer, err := milpactl.GetPrinter(cmd)
	dieIfError(err, "Error getting printer for result")
	milpaObj, err := api.Decode(reply.Body)
	dieIfError(err, "")
	err = printer.PrintObj(milpaObj, os.Stdout)
	if err != nil {
		// Just print the body of the response
		data, err2 := json.MarshalIndent(milpaObj, "", "    ")
		if err2 != nil {
			data = []byte(fmt.Sprintf("%#v", milpaObj))
		}
		fmt.Printf("Printing failed: %v\nObject:\n%s", err, string(data))
	}
}

func GetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "Get Resource",
		Long:  `Get resource specified by Name`,
		Run: func(cmd *cobra.Command, args []string) {
			get(cmd, args)
		},
	}
	// wide formatting doesn't work at this time.  Probably want to
	// get it working when we have individual pod status fields in the
	// PodStatus object
	cmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|table.")
	return cmd
}
