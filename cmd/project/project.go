package project

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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
		table.SetHeader([]string{"Project ID", "User ID", "Name", "Description", "Code", "Abi", "Yaml", "Schema", "Address", "Topic", "Start Block", "End Block", "Is Deleted", "Created At", "Updated At"})

		for _, project := range projects {
			table.Append([]string{
				project.ProjectId,
				project.UserId,
				project.Name,
				project.Description,
				project.Code,
				project.Abi,
				strings.ReplaceAll(project.Yaml, "\n", ""),
				strings.ReplaceAll(project.Schema, "\n", ""),
				project.Address,
				project.Topic,
				strconv.FormatInt(project.StartBlock, 10),
				strconv.FormatInt(project.EndBlock, 10),
				strconv.FormatBool(project.IsDeleted),
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
