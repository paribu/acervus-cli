package auth

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/paribu/acervus-cli/src/credential"
	"github.com/spf13/cobra"
)

var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Acervus Authentication",
	Long: `Acervus Authentication allows you to switch between multiple stored credentials. 
Choose the one you want to activate, and it will be set as your current credential.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		credentials, err := credential.LoadCredentials()
		if err != nil {
			return fmt.Errorf("error loading credentials: %s", err)
		}

		if len(credentials) == 0 {
			return errors.New("no credentials found, you need to login first")
		}

		if len(credentials) == 1 {
			err := selectCredential(credentials, 0)
			if err == nil {
				cmd.Printf("%s has been set as the current active account.\n", credentials[0].Email)
			}
			return err
		}

		selectedIndex := -1

		if email != "" {
			selectedIndex = selectCredentialWithEmail(credentials, email)
		} else {
			selectedIndex = chooseCredential(cmd, credentials)
		}

		if selectedIndex == -1 {
			return errors.New("selected credential not found")
		}

		err = selectCredential(credentials, selectedIndex)
		if err != nil {
			return fmt.Errorf("error selecting credential: %s", err)
		}

		cmd.Printf("%s has been set as the current active account.\n", credentials[selectedIndex].Email)

		return nil
	},
}

func init() {
	AuthCmd.AddCommand(loginCmd)
	AuthCmd.AddCommand(registerCmd)
	AuthCmd.AddCommand(resetPasswordCmd)
	AuthCmd.AddCommand(logoutCmd)
	AuthCmd.AddCommand(passwordRecoveryCmd)

	AuthCmd.Flags().StringVarP(&email, "email", "e", "", "Switch to a specific credential by email")
}

func selectCredentialWithEmail(credentials []*credential.Credential, email string) int {
	var selectedIndex int = -1
	for i, cred := range credentials {
		if cred.Email == email {
			selectedIndex = i
			break
		}
	}

	return selectedIndex
}

func selectCredential(credentials []*credential.Credential, index int) error {
	err := credential.SelectCredential(credentials[index].Email)
	if err != nil {
		return fmt.Errorf("error selecting credential: %s", err)
	}
	return nil
}

func listCredentials(cmd *cobra.Command, credentials []*credential.Credential) {
	cmd.Println("Choose an account:")
	for i, cred := range credentials {
		cmd.Printf("%d: %s\n", i, cred.Email)
	}
}

func chooseCredential(cmd *cobra.Command, credentials []*credential.Credential) int {
	listCredentials(cmd, credentials)
	scanner := bufio.NewScanner(os.Stdin)

	selectedIndex := -1
	for {
		cmd.Print("Enter the number of the account you want to use: ")
		scanner.Scan()
		input := scanner.Text()

		var err error
		selectedIndex, err = strconv.Atoi(input)
		if err == nil && selectedIndex >= 0 && selectedIndex < len(credentials) {
			break
		}

		fmt.Println("Invalid choice, please try again.")
	}

	return selectedIndex
}
