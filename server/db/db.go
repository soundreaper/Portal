package db

import (
	"fmt"

	"github.com/soundreaper/portal/config"
)

func GetMySQLConnectionString() string {
	// get db config from environment
	config := config.GetConfig()

	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DbUsername,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName)

	return dataBase
}
