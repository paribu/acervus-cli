package project

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/spf13/cobra"
)

var deleteProjectCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a project",
	Long:  "Delete the project with given ID.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api := api.NewProjectManagerAPI()

		err := api.DeleteProject(projectID)
		if err != nil {
			return fmt.Errorf("error while deleting project: %s", err)
		}

		cmd.Printf("Project with ID %s has been successfully deleted.\n", projectID)

		return nil
	},
}

func init() {
	deleteProjectCmd.Flags().StringVarP(&projectID, "id", "i", "", "ID of the project you want to delete")
	deleteProjectCmd.MarkFlagRequired("id")
}
