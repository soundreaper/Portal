package config

import (
	"os"
)

// AuthConfiguration represents the values needed for auth configuration
type AuthConfiguration struct {
	Auth0Domain string
	Auth0ID     string
}

// GetAuthConfig will return the default auth configuration
func GetAuthConfig() *AuthConfiguration {
	return &AuthConfiguration{
		Auth0Domain: os.Getenv("AUTH0_DOMAIN"),
		Auth0ID:     os.Getenv("AUTH0_API_ID"),
	}
}
