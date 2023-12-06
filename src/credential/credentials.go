package credential

import (
	"encoding/json"
	"errors"
	"os"
)

const credentialsFile = "./credentials.json"

// AddCredential adds new credential into credentials file.
// However there are different contexts and function behaves differently based on the given context.
// For "login" context we try to update existing account's tokens.
// For "register" context we try to add new records to the file.
func AddCredential(ctx Context, email, refreshToken, accessToken string) error {
	if !isCredentialsFileExists() {
		err := saveCredentials([]*Credential{})
		if err != nil {
			return err
		}
	}

	credentials, err := LoadCredentials()
	if err != nil {
		return err
	}

	for i, cred := range credentials {
		credentials[i].Current = false

		if cred.Email == email {
			if ctx == RegisterContext {
				return errors.New("credential with the given email already exists")
			}
			if ctx == LoginContext {
				credentials[i].RefreshToken = refreshToken
				credentials[i].AccessToken = accessToken
				credentials[i].Current = true

				return saveCredentials(credentials)
			}
		}
	}

	credentials = append(credentials, &Credential{
		Email:        email,
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		Current:      true,
	})

	return saveCredentials(credentials)
}

// UpdateCredential updates the record with given email from credentials file.
// It can also add new credentials (e.g. user login after credentials file deleted).
func UpdateCredential(email, refreshToken, accessToken string) error {
	credentials, err := LoadCredentials()
	if err != nil {
		return err
	}

	var found bool
	for i := range credentials {
		if credentials[i].Email == email {
			credentials[i].RefreshToken = refreshToken
			credentials[i].AccessToken = accessToken
			found = true
			break
		}
	}

	if !found {
		// RefreshContext is used for refreshing expired access tokens.
		return AddCredential(RefreshContext, email, refreshToken, accessToken)
	}

	return saveCredentials(credentials)
}

// RemoveCredential deletes the record with given email from credentials file.
func RemoveCredential(email string) error {
	credentials, err := LoadCredentials()
	if err != nil {
		return err
	}

	var found bool
	var newCredentials []*Credential
	for _, cred := range credentials {
		if cred.Email != email {
			newCredentials = append(newCredentials, cred)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("credential with the given email not found")
	}

	return saveCredentials(newCredentials)
}

// SelectCredential allows user to select currently active credential.
func SelectCredential(email string) error {
	credentials, err := LoadCredentials()
	if err != nil {
		return err
	}

	var found bool
	for i := range credentials {
		if credentials[i].Email == email {
			credentials[i].Current = true
			found = true
		} else {
			credentials[i].Current = false
		}
	}

	if !found {
		return errors.New("credential with the given email not found")
	}

	return saveCredentials(credentials)
}

// LoadCredentials reads existing credentials from credentials file.
func LoadCredentials() ([]*Credential, error) {
	if !isCredentialsFileExists() {
		return nil, errors.New("no credentials found")
	}

	data, err := os.ReadFile(credentialsFile)
	if err != nil {
		return nil, err
	}

	var credentials []*Credential
	err = json.Unmarshal(data, &credentials)
	if err != nil {
		return nil, err
	}

	return credentials, nil
}

// GetCurrentCredential returns currently active credential.
func GetCurrentCredential() (*Credential, error) {
	credentials, err := LoadCredentials()
	if err != nil {
		return nil, errors.New("no credentials found")
	}

	for _, cred := range credentials {
		if cred.Current {
			return cred, nil
		}
	}

	return nil, errors.New("no active credential found")
}

// isCredentialsFileExists checks if credentials file exists.
func isCredentialsFileExists() bool {
	_, err := os.Stat(credentialsFile)
	return err == nil
}

// saveCredentials writes given credentials into credentials file.
func saveCredentials(credentials []*Credential) error {
	data, err := json.MarshalIndent(credentials, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(credentialsFile, data, 0644)
}
