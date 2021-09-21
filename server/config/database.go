package config

import (
	"os"
)

// DBConfiguration represents the values needed for database configuration
type DBConfiguration struct {
	Username string
	Password string
	Port     string
	Host     string
	DBName   string
}

// GetDBConfig will return the default database connection
func GetDBConfig() *DBConfiguration {
	return &DBConfiguration{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
