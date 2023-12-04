package deploy

import (
	"github.com/spf13/cobra"
)

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a project into Acervus",
	Long: `The "deploy" command is used to initiate the deployment process of a project identified by its project ID into Acervus. 
This command interacts with the Acervus platform to start the deployment and reports the result.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO implement
		return nil
	},
}

func init() {
	DeployCmd.Flags().StringVarP(&settingsFilePath, "settings", "s", "./settings.yaml", "Settings YAML file")
	DeployCmd.Flags().StringVarP(&projectFilePath, "project", "p", "", "project.ts file")
	DeployCmd.Flags().StringVarP(&projectID, "id", "i", "", "ID of the project you want to deploy")
	DeployCmd.MarkFlagRequired("id")
}
