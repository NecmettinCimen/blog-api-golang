package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectPostgres() *gorm.DB {
	dsn := "user=username password=password dbname=default_database host=localhost sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
