package project

import "github.com/spf13/cobra"

var exportProjectCmd = &cobra.Command{
	Use:   "export",
	Short: "Export",
	Long:  `Export the results of a project`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO implement
		return nil
	},
}

func init() {
	exportProjectCmd.Flags().StringVarP(&projectID, "id", "p", "", "ID of the project you want to export")
	exportProjectCmd.MarkFlagRequired("id")
}
