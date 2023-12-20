package query

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/spf13/cobra"
)

var queryLogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "List user logs",
	Long:  "List user logs based on given filters",
	RunE: func(cmd *cobra.Command, args []string) error {
		pmApi := api.NewProjectManagerAPI()

		logsListStr, err := pmApi.ListProjectLog(api.ProjectLogListRequest{
			ProjectId: projectId,
			Level:     logLevel.String(),
			Limit:     limit,
			Page:      1, // implement pagination if needed
		})
		if err != nil {
			return fmt.Errorf("error while getting logs list: %s", err)
		}

		if pretty {
			// pretty print json string
			logsListStr, err = prettyJsonString(logsListStr)
			if err != nil {
				return fmt.Errorf("error while pretty printing logs list: %s", err)
			}
		} else {
			logsListStr = fmt.Sprintf("%s\n", logsListStr)
		}

		cmd.Println("Logs list:")
		cmd.Print(logsListStr)

		return nil
	},
}

func init() {
	queryLogsCmd.Flags().StringVarP(&projectId, "id", "i", "", "Project ID to filter query by")
	queryLogsCmd.Flags().VarP(
		&logLevel,
		"level",
		"v",
		fmt.Sprintf(`Log level to filter query by. Possible values: "%s", "%s"`, errorLogLevel, infoLogLevel),
	)
	queryLogsCmd.Flags().IntVarP(&limit, "limit", "l", 10, "Limit of entities to return")
	queryLogsCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty logs list output")
}
