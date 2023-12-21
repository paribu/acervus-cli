package project

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/spf13/cobra"
)

var ProjectCmd = &cobra.Command{
	Use:   "projects",
	Short: "List user projects",
	Long:  "List all the projects associated with the current user.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api := api.NewProjectManagerAPI()

		projects, err := api.ListProjects()
		if err != nil {
			return fmt.Errorf("error while getting projects list: %s", err)
		}

		for _, project := range projects {
			// TODO: pretty json data before printing
			cmd.Println(project)
		}

		return nil
	},
}

func init() {
	ProjectCmd.AddCommand(createProjectCmd)
	ProjectCmd.AddCommand(pauseProjectCmd)
	ProjectCmd.AddCommand(resumeProjectCmd)
	ProjectCmd.AddCommand(deleteProjectCmd)
	ProjectCmd.AddCommand(exportProjectCmd)
}
