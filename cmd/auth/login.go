package auth

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/paribu/acervus-cli/src/credential"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login into Acervus",
	Long: `Login to interact with Acervus cloud services. 
Once authenticated, your credentials will be stored locally.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Printf("Trying to login into Acervus with %s\n", email)

		api := api.NewAuthAPI()
		refreshToken, accessToken, err := api.Login(email, password)
		if err != nil {
			return fmt.Errorf("login failed: %s", err.Error())
		}

		err = credential.AddCredential(credential.LoginContext, email, refreshToken, accessToken)
		if err != nil {
			return fmt.Errorf("adding credential failed: %s", err.Error())
		}

		cmd.Println("Login successful")

		return nil
	},
}

func init() {
	loginCmd.Flags().StringVarP(&email, "email", "e", "", "Your email")
	loginCmd.MarkFlagRequired("email")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Your password")
	loginCmd.MarkFlagRequired("password")
}
