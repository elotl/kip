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
	"github.com/spf13/cobra"
)

func update(cmd *cobra.Command) {
	appManifestFile, err := cmd.Flags().GetString("file")
	dieIfError(err, "Error accessing 'file' flag for cmd %s", cmd.Name())
	client, conn, err := getKipClient(cmd.InheritedFlags(), true)
	dieIfError(err, "Failed to create kip client")
	defer conn.Close()

	errors := modify(client, appManifestFile, modifyUpdate)
	if len(errors) > 0 {
		fatal("Failed to update some resources: %v", errors)
	}
}

func UpdateCommand() *cobra.Command {

	var updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update kip object",
		Long:  `Update object specified in manifest on cloud of choice`,
		Run: func(cmd *cobra.Command, args []string) {
			update(cmd)
		},
	}
	updateCmd.Flags().StringP("file", "f", "", "Fully qualified path to manifest file")
	return updateCmd
}
