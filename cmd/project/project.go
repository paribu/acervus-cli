package project

import "github.com/spf13/cobra"

var ProjectCmd = &cobra.Command{
	Use:   "projects",
	Short: "List user projects",
	Long:  "List all the projects associated with the current user.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO implement
		return nil
	},
}

func init() {
	ProjectCmd.AddCommand(createProjectCmd)
	ProjectCmd.AddCommand(deleteProjectCmd)
	ProjectCmd.AddCommand(exportProjectCmd)
}
