package query

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/spf13/cobra"
)

var queryDataCmd = &cobra.Command{
	Use:   "data",
	Short: "List user data",
	Long:  "List user data based on given filters",
	RunE: func(cmd *cobra.Command, args []string) error {
		pmApi := api.NewProjectManagerAPI()

		dataListStr, err := pmApi.ListProjectData(api.ProjectDataListRequest{
			ProjectId: projectId,
			Name:      name,
			Value:     filtersStr,
			Limit:     limit,
			Page:      1, // implement pagination if needed
		})
		if err != nil {
			return fmt.Errorf("error while getting data list: %s", err)
		}

		if pretty {
			// pretty print json string
			dataListStr, err = prettyJsonString(dataListStr)
			if err != nil {
				return fmt.Errorf("error while pretty printing data list: %s", err)
			}
		} else {
			dataListStr = fmt.Sprintf("%s\n", dataListStr)
		}

		cmd.Println("Data list:")
		cmd.Print(dataListStr)

		return nil
	},
}

func prettyJsonString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func init() {
	queryDataCmd.Flags().StringVarP(&projectId, "id", "i", "", "Project ID to filter query by")
	queryDataCmd.Flags().StringVarP(&name, "name", "n", "", "Name of entity to filter query by")
	queryDataCmd.Flags().StringVarP(&filtersStr, "filters", "f", "", "Json string of filters to apply to query")
	queryDataCmd.Flags().IntVarP(&limit, "limit", "l", 10, "Limit of entities to return")
	queryDataCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty data list output")
}
