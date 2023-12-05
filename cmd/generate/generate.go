package generate

import "github.com/spf13/cobra"

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate",
	Long:  `Generate settings file from scratch or boilerplate from previously created settings file`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("You should select \"settings\", \"boilerplate\" or \"graphql\" command to continue generating.")
	},
}

func init() {
	GenerateCmd.AddCommand(generateBoilerplateCmd)
	GenerateCmd.AddCommand(generateGraphQLCmd)
	GenerateCmd.AddCommand(GenerateSettingsCmd)
}
