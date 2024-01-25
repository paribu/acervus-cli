package project

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/paribu/acervus-cli/src/api"
	"github.com/spf13/cobra"
)

var ProjectCmd = &cobra.Command{
	Use:   "projects",
	Short: "List user projects",
	Long:  "List all the projects associated with the current user.",
	RunE: func(cmd *cobra.Command, args []string) error {
		api := api.NewProjectManagerAPI()

		projects, err := api.ListProjects()
		if err != nil {
			return fmt.Errorf("error while getting projects list: %s", err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Project ID", "User ID", "Name", "Description", "Address", "Topic", "Start Block", "End Block", "Created At", "Updated At"})

		for _, project := range projects {
			table.Append([]string{
				project.ProjectId,
				project.UserId,
				project.Name,
				project.Description,
				project.Address,
				project.Topic,
				strconv.FormatInt(project.StartBlock, 10),
				strconv.FormatInt(project.EndBlock, 10),
				project.CreatedAt,
				project.UpdatedAt,
			})
		}

		table.Render()

		return nil
	},
}

func init() {
	ProjectCmd.AddCommand(createProjectCmd)
	ProjectCmd.AddCommand(pauseProjectCmd)
	ProjectCmd.AddCommand(resumeProjectCmd)
	ProjectCmd.AddCommand(deleteProjectCmd)
	ProjectCmd.AddCommand(exportProjectCmd)
}
