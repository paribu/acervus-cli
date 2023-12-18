package generate

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/spf13/cobra"
)

var generateGraphQLCmd = &cobra.Command{
	Use:   "graphql",
	Short: "Generate GraphQL",
	Long:  "Generate GraphQL schema from ABI and settings file",
	RunE: func(cmd *cobra.Command, args []string) error {
		api := api.NewProjectManagerAPI()

		response, err := api.GraphQL(gqlProjectDir, settingsFilePath, autoSkip)
		if err != nil {
			return fmt.Errorf("error when generating files: %s", err)
		}

		for _, fileInfo := range response.Files {
			fileName := filepath.Join(fileInfo.Path)
			fileDir := filepath.Dir(fileName)

			if _, err := os.Stat(fileDir); errors.Is(err, os.ErrNotExist) {
				os.MkdirAll(fileDir, os.ModePerm)
			}

			err = os.WriteFile(fileName, []byte(fileInfo.Contents), os.ModePerm)
			if err != nil {
				return fmt.Errorf("error when writing files: %s", err)
			}
		}

		return nil
	},
}

func init() {
	generateGraphQLCmd.Flags().StringVarP(&gqlProjectDir, "dir", "d", "", "Directory where graphql files will be created")
	generateGraphQLCmd.MarkFlagRequired("dir")

	generateGraphQLCmd.Flags().StringVarP(&settingsFilePath, "settings", "s", "./settings.yaml", "Path to settings file")
	generateGraphQLCmd.Flags().BoolVarP(&autoSkip, "auto-skip", "a", false, "Set skip mode")
}
