package auth

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/spf13/cobra"
)

var resetPasswordCmd = &cobra.Command{
	Use:   "reset-password",
	Short: "Reset your password",
	Long:  `Reset your password, by providing the code you received in your email.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Printf("Resetting password for %s...\n", email)

		api := api.NewAuthAPI()

		resp, err := api.ResetPassword(email, password, verificationCode)
		if err != nil {
			cmd.Println(resp)
			return fmt.Errorf("could not reset password: %v", err)
		}

		cmd.Println("Password reset successfully")

		return nil
	},
}

func init() {
	resetPasswordCmd.Flags().StringVarP(&email, "email", "e", "", "Your email")
	resetPasswordCmd.MarkFlagRequired("email")
	resetPasswordCmd.Flags().StringVarP(&password, "password", "p", "", "Your password")
	resetPasswordCmd.MarkFlagRequired("password")
	resetPasswordCmd.Flags().StringVarP(&verificationCode, "code", "c", "", "The code you received in your email")
	resetPasswordCmd.MarkFlagRequired("code")
}
