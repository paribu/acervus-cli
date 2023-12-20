package project

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/spf13/cobra"
)

var pauseProjectCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause a project",
	Long:  "Pause a project and stop all running services related to it temporarily",
	RunE: func(cmd *cobra.Command, args []string) error {
		api := api.NewProjectManagerAPI()

		err := api.PauseProject(projectID)
		if err != nil {
			return fmt.Errorf("error while pausing project: %s", err)
		}

		cmd.Printf("Project with ID %s has been successfully paused.\n", projectID)

		return nil
	},
}

func init() {
	pauseProjectCmd.Flags().StringVarP(&projectID, "id", "i", "", "ID of the project you want to pause")
	pauseProjectCmd.MarkFlagRequired("id")
}
