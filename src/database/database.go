package database

import (
	"os"
	_ "docker-air-echo/autoload"
	"docker-air-echo/constants"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func init() {
	var err error

	database, err = gorm.Open(postgres.Open(os.Getenv(constants.ENV_DB_URI)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	return database
}
