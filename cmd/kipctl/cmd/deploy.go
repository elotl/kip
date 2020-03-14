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
	"io"
	"os"

	"github.com/elotl/kip/pkg/clientapi"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

func deploy(cmd *cobra.Command, args []string) {
	resourceName := args[0]
	itemName := args[1]
	pkgfile := args[2]

	client, conn, err := getKipClient(cmd.InheritedFlags(), false)
	dieIfError(err, "Failed to create kip client")
	defer conn.Close()

	req := &clientapi.DeployRequest{
		ResourceName: resourceName,
		ItemName:     itemName,
	}
	f, err := os.Open(pkgfile)
	dieIfError(err, "Could not open package file %s", pkgfile)
	stream, err := client.Deploy(context.Background())
	dieIfError(err,
		"Could not deploy %s for %s/%s", pkgfile, resourceName, itemName)
	for {
		buf := make([]byte, 64*1024) // Recommended chunk size for streaming.
		_, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		dieIfError(err, "Could not read package file %s", pkgfile)
		req.PackageData = buf
		err = stream.Send(req)
		dieIfError(err, "Could not send package data")
	}
	reply, err := stream.CloseAndRecv()
	dieIfError(err,
		"Could not deploy %s for %s/%s", pkgfile, resourceName, itemName)
	dieIfReplyError("Deploy", reply)
}

func DeployCommand() *cobra.Command {
	var deployCmd = &cobra.Command{
		Use:   "deploy pod_name package_name package_file",
		Short: "Deploy Kip package for a pod",
		Long:  `Deploy Kip package for a pod`,
		Args:  cobra.RangeArgs(3, 3),
		Run:   deploy,
	}
	return deployCmd
}
