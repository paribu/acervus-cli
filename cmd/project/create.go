package project

import (
	"github.com/spf13/cobra"
)

var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project",
	Long:  "Create a new project with default files.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO implement
		return nil
	},
}

func init() {
	createProjectCmd.Flags().StringVarP(&projectDir, "dir", "d", "", "Directory where the project will be created")
}
