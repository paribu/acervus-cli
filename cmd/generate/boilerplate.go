package generate

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/paribu/acervus-cli/src/prompt"
	"github.com/spf13/cobra"
)

var generateBoilerplateCmd = &cobra.Command{
	Use:   "boilerplate",
	Short: "Generates boilerplate code for your project.",
	Long: `This command generates boilerplate code for your project.
It's automatically runs when you create a new project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		api := api.NewProjectManagerAPI()

		if projectDir == "" {
			projectDir = prompt.GetProjectDirectory()
		}

		boilerplateRes, err := api.GenerateBoilerplate(settingsFilePath)
		if err != nil {
			return fmt.Errorf("error when generating files: %s", err)
		}

		for _, fileInfo := range boilerplateRes.Files {
			fileName := filepath.Join(projectDir, fileInfo.Path)
			fileDir := filepath.Dir(fileName)

			if _, err := os.Stat(fileDir); errors.Is(err, os.ErrNotExist) {
				os.MkdirAll(fileDir, os.ModePerm)
			}

			if err = os.WriteFile(fileName, []byte(fileInfo.Contents), os.ModePerm); err != nil {
				return fmt.Errorf("error when writing files: %s", err)
			}
		}

		return nil
	},
}

func init() {
	generateBoilerplateCmd.Flags().StringVarP(&projectDir, "dir", "d", "", "Directory where the generated files will be placed")
	generateGraphQLCmd.Flags().StringVarP(&settingsFilePath, "settings", "s", "./settings.yaml", "Path to settings file")
}
