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
