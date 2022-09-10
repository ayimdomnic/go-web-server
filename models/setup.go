package models

import (
	"github.com/ayimdomnic/go-web-server/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := helpers.GoDotEnvVariableKey("DSN")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&User{}, &Lotto{}, &Ticket{}, &Account{}, &Transaction{})

	DB = db
}
