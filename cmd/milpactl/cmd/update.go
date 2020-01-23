package cmd

import (
	"github.com/spf13/cobra"
)

func update(cmd *cobra.Command, args []string) {
	appManifestFile, err := cmd.Flags().GetString("file")
	dieIfError(err, "Error accessing 'file' flag for cmd %s", cmd.Name())
	client, conn, err := getMilpaClient(cmd.InheritedFlags(), true)
	dieIfError(err, "Failed to create milpa client")
	defer conn.Close()

	errors := modify(client, appManifestFile, modifyUpdate)
	if len(errors) > 0 {
		fatal("Failed to update some resources: %v", errors)
	}
}

func UpdateCommand() *cobra.Command {

	var updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update milpa object",
		Long:  `Update object specified in manifest on cloud of choice`,
		Run: func(cmd *cobra.Command, args []string) {
			update(cmd, args)
		},
	}
	updateCmd.Flags().StringP("file", "f", "", "Fully qualified path to manifest file")
	return updateCmd
}
