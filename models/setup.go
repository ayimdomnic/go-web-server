package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "host=ec2-34-231-42-166.compute-1.amazonaws.com user=gfaxtykdofxihj password=8b00c2e63a669c1b5d11bffbf7d3ad39ed68a8b3f34fcfc5d3aea2a3ab7aa5af dbname=d4pbcfl1obr6uo port=5432 sslmode=true TimeZone=Africa/Nairobi"
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
