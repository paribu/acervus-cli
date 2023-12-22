package project

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/spf13/cobra"
)

var exportProjectCmd = &cobra.Command{
	Use:   "export",
	Short: "Export",
	Long:  `Export results of a project`,
	RunE: func(cmd *cobra.Command, args []string) error {
		api := api.NewProjectManagerAPI()

		err := api.ExportProject(projectID)
		if err != nil {
			return fmt.Errorf("error while exporting project: %s", err)
		}

		cmd.Printf("Project with ID %s has been successfully exported.\n", projectID)

		return nil
	},
}

func init() {
	exportProjectCmd.Flags().StringVarP(&projectID, "id", "i", "", "ID of the project you want to export")
	exportProjectCmd.MarkFlagRequired("id")
}
