package migrate

import "github.com/spf13/cobra"

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate your project from a specified source to the current platform.",
	Long: `The migrate command enables users to effortlessly transfer their projects
from a designated source platform to the current system.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// TODO implement
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO implement
		return nil
	},
}

func init() {
	MigrateCmd.Flags().StringVarP(
		&sourcePlatform,
		"sourcePlatform",
		"s",
		"",
		"Select the platform you will be migrating from",
	)
	MigrateCmd.Flags().StringVarP(
		&projectDir,
		"projectDir",
		"d",
		"",
		"Directory where the project will be created",
	)
}
