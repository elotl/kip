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

package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/clientapi"
	"github.com/elotl/kip/pkg/kipctl"
	"github.com/elotl/kip/pkg/util"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

func get(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		msg := "Not enough arguments\nUsage: kipctl get (pod|node|service|resource) [Name]"
		fatal(msg)
	}
	kind := kipctl.CleanupResourceName(args[0])
	var name string
	if len(args) > 1 {
		name = args[1]
	}
	if !util.StringInSlice(kind, getTypes) {
		fatal("Illegal resource type for GET: %s", kind)
	}

	client, conn, err := getKipClient(cmd.InheritedFlags(), false)
	dieIfError(err, "Failed to create kip client")
	defer conn.Close()

	getRequest := &clientapi.GetRequest{
		Kind: []byte(kind),
		Name: []byte(name),
	}
	reply, err := client.Get(context.Background(), getRequest)
	dieIfError(err, "Could not get resource")
	dieIfReplyError("Get", reply)
	printer, err := kipctl.GetPrinter(cmd)
	dieIfError(err, "Error getting printer for result")
	kipObj, err := api.Decode(reply.Body)
	dieIfError(err, "")
	err = printer.PrintObj(kipObj, os.Stdout)
	if err != nil {
		// Just print the body of the response
		data, err2 := json.MarshalIndent(kipObj, "", "    ")
		if err2 != nil {
			data = []byte(fmt.Sprintf("%#v", kipObj))
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
