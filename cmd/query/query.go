package query

import (
	"github.com/spf13/cobra"
)

var QueryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query data/logs",
	Long:  "Query data/logs list based on given filters",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("You should select \"data\" or \"logs\" command to continue.")
	},
}

func init() {
	QueryCmd.AddCommand(queryDataCmd)
	// TODO: add logs command
}
