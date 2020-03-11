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
	"context"
	"fmt"
	"os"

	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/spf13/cobra"
)

func dump(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Not enough arguments")
		fatal("Usage: kipctl dump <all|stack|PodController|NodeController>")
	}
	controller := args[0]
	client, conn, err := getKipClient(cmd.InheritedFlags(), true)
	dieIfError(err, "Failed to create kip client")
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
