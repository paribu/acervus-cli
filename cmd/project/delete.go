package project

import (
	"github.com/spf13/cobra"
)

var deleteProjectCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a project",
	Long:  "Delete the project with given ID.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO implement
		return nil
	},
}

func init() {
	deleteProjectCmd.Flags().StringVarP(&projectID, "id", "p", "", "ID of the project you want to delete")
	deleteProjectCmd.MarkFlagRequired("id")
}
