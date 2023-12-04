package auth

import (
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register with Acervus",
	Long: `Register on Acervus platform using an email and password. 
The credentials will then be stored locally.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO implement
		return nil
	},
}

func init() {
	registerCmd.Flags().StringVarP(&email, "email", "e", "", "Your email")
	registerCmd.MarkFlagRequired("email")
	registerCmd.Flags().StringVarP(&password, "password", "p", "", "Your password")
	registerCmd.MarkFlagRequired("password")
}
