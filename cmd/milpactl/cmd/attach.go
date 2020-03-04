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
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/wsstream"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	attachPodName     string
	attachUnitName    string
	attachInteractive bool
	attachUsageStr    = "attach POD_NAME"
)

func attach(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fatal("A pod name is required: " + attachUsageStr)
	}
	attachPodName = args[0]

	params := api.AttachParams{
		PodName:     attachPodName,
		UnitName:    attachUnitName,
		Interactive: attachInteractive,
		TTY:         false,
	}

	client, conn, err := getMilpaClient(cmd.InheritedFlags(), false)
	dieIfError(err, "Failed to create milpa client")
	defer conn.Close()

	stream, err := client.Attach(context.Background())
	dieIfError(err, "Failed to setup attach streaming client")

	b, err := json.Marshal(params)
	dieIfError(err, "Error serializing attach parameters")
	paramMsg := &clientapi.StreamMsg{Data: b}
	err = stream.Send(paramMsg)
	dieIfError(err, "Error sending initial attach parameters")

	// this looks a lot like exec. The count is at two.
	// https://en.wikipedia.org/wiki/Rule_of_three_(computer_programming)

	// Read from local stdin
	go func() {
		defer stream.CloseSend()
		// We read based on newlines. Using scanner won't work for
		// interactive programs but lets not worry about that now.
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			t := scanner.Text() + "\n"
			f := wsstream.PackMessage(wsstream.StdinChan, []byte(t))
			sm := &clientapi.StreamMsg{Data: f}
			if err := stream.Send(sm); err != nil {
				return
			}
		}
		dieIfError(scanner.Err(), "Error reading stdin")
	}()

	// Write to local stdout and stderr, if we get an exit code,
	// exit with that code
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return
		}
		dieIfError(err, "Error in grpc receive")
		c, msg, err := wsstream.UnpackMessage(resp.Data)
		dieIfError(err, "error unserializing websocket data")
		if len(msg) > 0 {
			if c == wsstream.StdoutChan {
				fmt.Fprint(os.Stdout, string(msg))
			} else if c == wsstream.StderrChan {
				fmt.Fprint(os.Stderr, string(msg))
			} else if c == wsstream.ExitCodeChan {
				i, err := strconv.Atoi(string(msg))
				if err != nil {
					errmsg := fmt.Sprintf("Invalid exit code: %s", msg)
					fmt.Fprint(os.Stderr, errmsg)
					os.Exit(1)
				}
				os.Exit(i)
			}
		}
	}
}

func AttachCommand() *cobra.Command {
	var attachCmd = &cobra.Command{
		Use:   "attach",
		Short: "Attach to a process that is already running inside an existing unit",
		Long:  `Attach to a process that is already running inside an existing unit`,
		Example: `# Get output from running pod my-pod, using the first unit by default
milpactl attach my-pod

# Get output from rubyserver unit from pod my-pod
milpactl attach my-pod -u rubyserver

# Send stdin to rubyserver in pod my-pod and sends stdout/stderr from rubyserver back to the client
milpactl attach my-pod -u rubyserver -i`,
		Run: func(cmd *cobra.Command, args []string) {
			attach(cmd, args)
		},
	}

	attachCmd.Flags().StringVarP(&attachUnitName, "unit", "u", "", "Unit name. If empty the first unit in the pod will be used")
	attachCmd.Flags().BoolVarP(&attachInteractive, "stdin", "i", false, "Pass stdin to the unit")

	return attachCmd
}
