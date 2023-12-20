package project

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/spf13/cobra"
)

var resumeProjectCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resume a project",
	Long:  "Resume a project and stop all running services related to it temporarily",
	RunE: func(cmd *cobra.Command, args []string) error {
		api := api.NewProjectManagerAPI()

		err := api.ResumeProject(projectID)
		if err != nil {
			return fmt.Errorf("error while pausing project: %s", err)
		}

		cmd.Printf("Project with ID %s has been successfully resumed.\n", projectID)

		return nil
	},
}

func init() {
	resumeProjectCmd.Flags().StringVarP(&projectID, "id", "i", "", "ID of the project you want to resume")
	resumeProjectCmd.MarkFlagRequired("id")
}
