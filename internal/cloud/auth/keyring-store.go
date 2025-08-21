package auth

import (
	"fmt"
	"strconv"

	"github.com/zalando/go-keyring"
)

func storeTokenInMemory(tokens TokenResponse) error {
	err := keyring.Set(serviceName, "access_token", tokens.AccessToken)
	if err != nil {
		return fmt.Errorf("unable to store access_token: %w", err)
	}

	err = keyring.Set(serviceName, "id_token", tokens.IDToken)
	if err != nil {
		return fmt.Errorf("unable to store id_token: %w", err)
	}

	err = keyring.Set(serviceName, "refresh_token", tokens.RefreshToken)
	if err != nil {
		return fmt.Errorf("unable to store refresh_token: %w", err)
	}

	err = keyring.Set(serviceName, "expires_in", strconv.Itoa(tokens.ExpiresIn))
	if err != nil {
		return fmt.Errorf("unable to store refresh_token: %w", err)
	}

	err = keyring.Set(serviceName, "token_type", tokens.TokenType)
	if err != nil {
		return fmt.Errorf("unable to store token_type: %w", err)
	}

	return nil
}

func retriveTokenFromMemory() (*TokenResponse, error) {
	var tokens TokenResponse

	accessToken, err := keyring.Get(serviceName, "access_token")
	if err != nil {
		return nil, fmt.Errorf("token not found")
	}

	tokens.AccessToken = accessToken

	idToken, err := keyring.Get(serviceName, "id_token")
	if err != nil {
		return nil, fmt.Errorf("token not found")
	}

	tokens.IDToken = idToken

	refreshToken, err := keyring.Get(serviceName, "refresh_token")
	if err != nil {
		return nil, fmt.Errorf("token not found")
	}

	tokens.RefreshToken = refreshToken

	tokenType, err := keyring.Get(serviceName, "token_type")
	if err != nil {
		return nil, fmt.Errorf("token not found")
	}

	tokens.TokenType = tokenType

	return &tokens, nil
}

func clearTokenInMemory() {
	keyring.DeleteAll(serviceName)
}

func Reset() {
	clearTokenInMemory()
}
