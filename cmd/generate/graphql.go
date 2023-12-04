package generate

import "github.com/spf13/cobra"

var generateGraphQLCmd = &cobra.Command{
	Use:   "graphql",
	Short: "Generate GraphQL",
	Long:  "Generate GraphQL schema from ABI and settings file",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO implement
		return nil
	},
}

func init() {
	generateGraphQLCmd.Flags().StringVarP(&settingsFilePath, "settings", "s", "./settings.yaml", "Path to settings file")
}
