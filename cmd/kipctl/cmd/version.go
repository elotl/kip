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
	"fmt"

	"github.com/elotl/kip/pkg/clientapi"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

func getServerVersion(cmd *cobra.Command, args []string) {
	client, conn, err := getKipClient(cmd.InheritedFlags(), false)
	dieIfError(err, "Failed to create kip client")
	defer conn.Close()
	versionRequest := &clientapi.VersionRequest{}
	reply, err := client.GetVersion(context.Background(), versionRequest)
	dieIfError(err, "could not gather grpc server version")
	fmt.Printf("Kip server version: %s\n", string(reply.VersionInfo))
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
