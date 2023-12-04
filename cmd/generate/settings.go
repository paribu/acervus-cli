package generate

import "github.com/spf13/cobra"

var generateSettingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Generates a settings template file.",
	Long: `This command generates a new settings file with a default template.
This file will serve as a configuration for your application, outlining all the
necessary settings. Users can then customize the settings to fit their requirements.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// TODO implement
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO implement
		return nil
	},
}
