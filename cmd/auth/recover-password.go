package auth

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/spf13/cobra"
)

var passwordRecoveryCmd = &cobra.Command{
	Use:   "recover-password",
	Short: "Initiate password recovery",
	Long:  `Request a password recovery token for the specified email.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Printf("Initiating password recovery for %s\n", email)

		api := api.NewAuthAPI()

		resp, err := api.RecoverPassword(email)
		if err != nil {
			fmt.Println(resp)
			return err
		}

		fmt.Println("Password recovery email sent. Check your inbox for instructions.")
		return nil
	},
}

func init() {
	passwordRecoveryCmd.Flags().StringVarP(&email, "email", "e", "", "Your email")
	passwordRecoveryCmd.MarkFlagRequired("email")
}
