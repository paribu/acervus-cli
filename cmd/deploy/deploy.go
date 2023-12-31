package deploy

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/spf13/cobra"
)

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a project into Acervus",
	Long: `The "deploy" command is used to initiate the deployment process of a project identified by its project ID into Acervus. 
This command interacts with the Acervus platform to start the deployment and reports the result.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Printf("Project %s trying to deploy into Acervus \n", projectID)

		api := api.NewProjectManagerAPI()

		_, err := api.Deploy(projectID, settingsFilePath, projectFilePath)
		if err != nil {
			return fmt.Errorf("could not deploy into Acervus: %v", err)
		}

		cmd.Println("Deploy successful")

		return nil
	},
}

func init() {
	DeployCmd.Flags().StringVarP(&settingsFilePath, "settings", "s", "./settings.yaml", "Path to settings file")
	DeployCmd.Flags().StringVarP(&projectFilePath, "project", "p", "", "Path to project.ts file")
	DeployCmd.Flags().StringVarP(&projectID, "id", "i", "", "ID of the project you want to deploy")
	DeployCmd.MarkFlagRequired("id")
}
