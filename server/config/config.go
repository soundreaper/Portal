package config

import (
	"os"

	"github.com/joho/godotenv"

	"log"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Did not load variables from .env file. This is normal for CI/CD or production.")
	}
}

type Config struct {
	AppEnv     string // the environment that the application is running in (env, prod, etc)
	DbUsername string // database username
	DbPassword string // database password
	DbHost     string // database host
	DbPort     string // databse port
}

// GetConfig will return the current config
func GetConfig() *Config {
	config := &Config{
		AppEnv:     os.Getenv("APP_ENV"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbUsername: os.Getenv("DB_USER"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
	}

	return config
}
