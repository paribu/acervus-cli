package test

import (
	"encoding/json"
	"fmt"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/spf13/cobra"
)

var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "Test",
	Long: `The "test" command is used to initiate testing for a Acervus Project. 
It allows you to validate and verify the functionality of the project using various testing scenarios and inputs.
This command initiates the testing process for the specified Acervus Project, providing detailed feedback on the success or failure of the tests.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Printf("Project %s is being tested in Acervus \n", projectID)

		api := api.NewProjectManagerAPI()

		testResponse, err := api.Test(projectID, settingsFilePath, projectFilePath)
		if err != nil {
			return fmt.Errorf("could not test in Acervus: %v", err)
		}

		testJSON, err := json.Marshal(testResponse)
		if err != nil {
			return fmt.Errorf("could not convert test results to JSON %v", err)
		}

		cmd.Println("Test successful")
		cmd.Println(string(testJSON))

		return nil
	},
}

func init() {
	TestCmd.Flags().StringVarP(&settingsFilePath, "settings", "s", "./settings.yaml", "Settings YAML file")
	TestCmd.Flags().StringVarP(&projectFilePath, "project", "p", "", "project.ts file")
	TestCmd.Flags().StringVarP(&projectID, "id", "i", "", "ID of the project you want to deploy")
	TestCmd.MarkFlagRequired("id")
}
