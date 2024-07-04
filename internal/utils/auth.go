package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const tokenFileName = ".cli_auth_tokens.json"

type AuthToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// SaveToken сохраняет токен аутентификации в файл
func SaveToken(ar AuthToken) error {
	tokenFile := filepath.Join(os.Getenv("HOME"), tokenFileName)
	file, err := os.Create(tokenFile)
	if err != nil {
		return err
	}
	defer CloseFile(file)

	if err := json.NewEncoder(file).Encode(ar); err != nil {
		return err
	}
	return nil
}

// GetToken загружает токен аутентификации из файла
func GetToken() string {
	tokens, err := LoadTokens()
	if err != nil {
		fmt.Printf("Error: failed %v\n", err)
		return ""
	}
	return tokens.AccessToken
}

// LoadTokens загружает токены аутентификации из файла
func LoadTokens() (*AuthToken, error) {
	tokenFile := filepath.Join(os.Getenv("HOME"), tokenFileName)
	file, err := os.Open(tokenFile)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file)

	var tokens AuthToken
	if err := json.NewDecoder(file).Decode(&tokens); err != nil {
		return nil, err
	}

	return &tokens, nil
}
