package auth

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/paribu/acervus-cli/src/credential"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of Acervus",
	Long:  `This command logs you out of Acervus and removes your credential from the local credentials file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		api := api.NewAuthAPI()

		err := api.Logout()
		if err != nil {
			return fmt.Errorf("an error occurred while logging out: %s", err.Error())
		}

		err = credential.RemoveCredential(email)
		if err != nil {
			return fmt.Errorf("an error occurred while logging out: %s", err.Error())
		}

		cmd.Println("Successfully logged out.")

		return nil
	},
}

func init() {
	logoutCmd.Flags().StringVarP(&email, "email", "e", "", "Your email")
	logoutCmd.MarkFlagRequired("email")
}
