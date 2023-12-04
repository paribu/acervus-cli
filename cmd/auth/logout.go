package auth

import (
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of Acervus",
	Long:  `This command logs you out of Acervus and removes your credential from the local credentials file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO implement
		return nil
	},
}

func init() {
	logoutCmd.Flags().StringVarP(&email, "email", "e", "", "Your email")
	logoutCmd.MarkFlagRequired("email")
}
