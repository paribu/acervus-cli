package generate

import "github.com/spf13/cobra"

var generateBoilerplateCmd = &cobra.Command{
	Use:   "boilerplate",
	Short: "Generates boilerplate code for your project.",
	Long: `This command generates boilerplate code for your project.
It's automatically runs when you create a new project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO implement
		return nil
	},
}

func init() {
	generateBoilerplateCmd.Flags().StringVarP(
		&projectDir,
		"projectDir",
		"d",
		"",
		"Directory where the generated files will be placed",
	)
}
