package cmd

import (
	"github.com/spf13/cobra"
)

func create(cmd *cobra.Command, args []string) {
	appManifestFile, err := cmd.Flags().GetString("file")
	dieIfError(err, "Error accessing 'file' flag for cmd %s", cmd.Name())
	client, conn, err := getMilpaClient(cmd.InheritedFlags(), true)
	dieIfError(err, "Failed to create milpa client")
	defer conn.Close()
	errors := modify(client, appManifestFile, modifyCreate)
	if len(errors) > 0 {
		fatal("Failed to create some resources: %v", errors)
	}
}

func CreateCommand() *cobra.Command {
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create milpa object",
		Long:  `Create object specified in manifest on cloud of choice`,
		Run: func(cmd *cobra.Command, args []string) {
			create(cmd, args)
		},
	}
	createCmd.Flags().StringP("file", "f", "", "Fully qualified path to manifest file")
	return createCmd
}
