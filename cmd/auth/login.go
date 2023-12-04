package auth

import (
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login into Acervus",
	Long: `Login to interact with Acervus cloud services. 
Once authenticated, your credentials will be stored locally.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO implement
		return nil
	},
}

func init() {
	loginCmd.Flags().StringVarP(&email, "email", "e", "", "Your email")
	loginCmd.MarkFlagRequired("email")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Your password")
	loginCmd.MarkFlagRequired("password")
}
