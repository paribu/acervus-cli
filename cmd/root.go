package cmd

import (
	"os"

	"github.com/paribu/acervus-cli/cmd/auth"
	"github.com/paribu/acervus-cli/cmd/deploy"
	"github.com/paribu/acervus-cli/cmd/generate"
	"github.com/paribu/acervus-cli/cmd/migrate"
	"github.com/paribu/acervus-cli/cmd/project"
	"github.com/paribu/acervus-cli/cmd/query"
	"github.com/paribu/acervus-cli/cmd/test"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "acervus",
	Version: "0.0.5",
	Short:   "CLI application to interact with Acervus Cloud",
	Long:    "You can use this application to manage your Acervus account, generate, test and deploy Acervus projects.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(auth.AuthCmd)
	rootCmd.AddCommand(deploy.DeployCmd)
	rootCmd.AddCommand(generate.GenerateCmd)
	rootCmd.AddCommand(migrate.MigrateCmd)
	rootCmd.AddCommand(project.ProjectCmd)
	rootCmd.AddCommand(test.TestCmd)
	rootCmd.AddCommand(query.QueryCmd)
}
