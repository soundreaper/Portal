package config

import (
	"os"
)

// AuthConfiguration represents the values needed for auth configuration
type AuthConfiguration struct {
	GoogleKey    string
	GoogleSecret string
}

// GetAuthConfig will return the default auth configuration
func GetAuthConfig() *AuthConfiguration {
	return &AuthConfiguration{
		GoogleKey:    os.Getenv("GOOGLE_KEY"),
		GoogleSecret: os.Getenv("GOOGLE_SECRET"),
	}
}
