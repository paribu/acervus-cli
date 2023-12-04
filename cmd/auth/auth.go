package auth

import "github.com/spf13/cobra"

var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Acervus Authentication",
	Long: `Acervus Authentication allows you to switch between multiple stored credentials. 
Choose the one you want to activate and it will be set as your current credential.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO implement
		return nil
	},
}

func init() {
	AuthCmd.AddCommand(loginCmd)
	AuthCmd.AddCommand(registerCmd)
	AuthCmd.AddCommand(logoutCmd)
}
