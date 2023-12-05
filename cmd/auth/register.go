package auth

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/api"
	"github.com/paribu/acervus-cli/src/credential"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register with Acervus",
	Long: `Register on Acervus platform using an email and password. 
The credentials will then be stored locally.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		credentials, _ := credential.LoadCredentials()
		for _, cred := range credentials {
			if cred.Email == email {
				return fmt.Errorf("there is already a user registered with the email %s", email)
			}
		}

		cmd.Printf("Trying to register into Acervus with %s\n", email)

		api := api.NewAuthAPI()

		refreshToken, accessToken, err := api.Register(email, password)
		if err != nil {
			return fmt.Errorf("could not register with Acervus: %v", err)
		}

		err = credential.AddCredential(credential.RegisterContext, email, refreshToken, accessToken)
		if err != nil {
			return fmt.Errorf("could not save credentials: %v", err)
		}

		cmd.Println("Registration successful. You are also automatically logged in and can start using your account")

		return nil
	},
}

func init() {
	registerCmd.Flags().StringVarP(&email, "email", "e", "", "Your email")
	registerCmd.MarkFlagRequired("email")
	registerCmd.Flags().StringVarP(&password, "password", "p", "", "Your password")
	registerCmd.MarkFlagRequired("password")
}
