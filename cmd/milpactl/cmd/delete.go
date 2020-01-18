package cmd

import (
	"fmt"

	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/milpactl"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

func del(cmd *cobra.Command, args []string) {
	// see if app manifest file has been supplied
	if len(args) > 0 && len(args) != 2 {
		fatal("Usage: milpactl delete <resource> <name>")
	}
	cascade, _ := cmd.Flags().GetBool("cascade")

	client, conn, err := getMilpaClient(cmd.InheritedFlags(), true)
	dieIfError(err, "Failed to create milpa client")
	defer conn.Close()

	if len(args) == 2 {
		kind := milpactl.CleanupResourceName(args[0])
		name := args[1]
		if !util.StringInSlice(kind, deleteTypes) {
			fatal("Illegal resource type: %s", kind)
		}
		deleteRequest := &clientapi.DeleteRequest{
			Kind:    []byte(kind),
			Name:    []byte(name),
			Cascade: cascade,
		}
		reply, err := client.Delete(context.Background(), deleteRequest)
		dieIfError(err, "Could not delete resource")
		dieIfReplyError("Delete", reply)
		fmt.Printf("%s\n", name)
	} else {
		manifestFile, err := cmd.Flags().GetString("file")
		dieIfError(err, "Error accessing 'file' flag for cmd %s", cmd.Name())
		op := modifyDeleteCascade
		if !cascade {
			op = modifyDelete
		}
		errors := modify(client, manifestFile, op)
		if len(errors) > 0 {
			fatal("Failed to update some resources: %v", errors)
		}
	}
}

func DeleteCommand() *cobra.Command {
	var deleteCmd = &cobra.Command{
		Use:   "delete ([-f filename] | (<resource> <name>))",
		Short: "Delete resource by filename or by resource and name",
		Long:  `Delete resource by filename or by resource and name`,
		Example: `
# Delete a pod using the type and name specified in the file pod.yml.
milpactl delete -f ./pod.yml

# Delete a pod named mypod
milpactl delete pod mypod

# Delete a deployment named mydeployment and delete all objects managed by that deployment
milpactl delete --cascade deployment mypod`,
		Run: func(cmd *cobra.Command, args []string) {
			del(cmd, args)
		},
	}
	deleteCmd.Flags().BoolP("cascade", "", true, "If true, cascade the deletion of the resources managed by this resource")
	deleteCmd.Flags().StringP("file", "f", "", "Fully qualified path to manifest file")
	return deleteCmd
}
