package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

const passwordsPath = "../data/passwords.json"

type Storage struct {
	credentials map[string][]credential
}

type credential struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (storage *Storage) SaveCredentials() error {
	data, err := json.MarshalIndent(storage.credentials, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(passwordsPath, data, 0600)
}

func (storage *Storage) loginExists(source, login string) bool {
	for _, credential := range storage.credentials[source] {
		if credential.Login == login {
			return true
		}
	}
	return false
}

func (storage *Storage) AddCredentials(source, login, password string) string {
	var err error
	if storage.credentials[source] == nil {
		storage.credentials[source] = make([]credential, 1)
		storage.credentials[source][0] = credential{login, password}
		err = storage.SaveCredentials()
	} else if storage.loginExists(source, login) {
		return "the credentials for the source are already exist"
	} else {
		storage.credentials[source] = append(storage.credentials[source], credential{login, password})
		err = storage.SaveCredentials()
	}
	if err == nil {
		return "saved"
	}
	return fmt.Sprintf("cannot save credentials due to this: %s", err)
}

func (storage *Storage) GetSources() string {
	var result strings.Builder
	result.WriteString("Credentials:\n")
	for source, credentials := range storage.credentials {
		fmt.Fprintln(&result, getCredentialsRaw(source, credentials, false))
	}
	return result.String()
}

func (storage *Storage) GetCredentials(source string) string {
	credentials := storage.credentials[source]
	if credentials == nil {
		return "error: source hasn't been found"
	}
	return getCredentialsRaw(source, credentials, true)
}

func (storage *Storage) GetPassword(source, login string) string {
	credentials := storage.credentials[source]
	if credentials == nil {
		return "error: source hasn't been found"
	}
	for _, credential := range credentials {
		if credential.Login == login {
			return getPasswordRaw(source, login, credential.Password)
		}
	}
	return "error: login hasn't been found"
}

func getCredentialsRaw(source string, credentials []credential, showPasswords bool) string {
	if showPasswords && len(credentials) == 1 {
		return getPasswordRaw(source, credentials[0].Login, credentials[0].Password)
	}
	return fmt.Sprintf("  Source: %s\n  Logins: %s\n", source, strings.Join(getLogins(credentials), ", "))
}

func getPasswordRaw(source, login, password string) string {
	return fmt.Sprintf("  Source: %s\n  Login: %s\n  Password: %s", source, login, password)
}

func getLogins(credentials []credential) []string {
	logins := make([]string, len(credentials))
	for i, credential := range credentials {
		logins[i] = credential.Login
	}
	return logins
}

func (storage *Storage) String() string {
	var result strings.Builder
	result.WriteString("Credentials:\n")
	for source, credentials := range storage.credentials {
		fmt.Fprintf(&result, "  Source: %s\n", source)
		for _, credential := range credentials {
			fmt.Fprintf(&result, "    Login: %s\n", credential.Login)
			fmt.Fprintf(&result, "    Password: %s\n", credential.Password)
			fmt.Fprint(&result, "\n")
		}
	}
	return result.String()
}

func LoadStorage() (*Storage, error) {
	storage := Storage{}
	var err error
	err = LoadCredentials(&storage)

	return &storage, err
}

func LoadCredentials(storage *Storage) error {
	data, err := os.ReadFile(passwordsPath)
	if err == nil {
		return json.Unmarshal(data, &storage.credentials)
	} else if os.IsNotExist(err) {
		if err := os.Mkdir("../data", 0700); err != nil && !errors.Is(err, fs.ErrExist) {
			return err
		}
		storage.credentials = make(map[string][]credential)
		if data, err = json.Marshal(storage.credentials); err != nil {
			return err
		}
		return os.WriteFile("../data/passwords.json", data, 0600)
	} else {
		return err
	}
}
