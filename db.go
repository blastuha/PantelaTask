package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	var DB *gorm.DB

	dsn := "host=localhost user=postgres password=aZAz1998 dbname=postgres port=5432 sslmode=disable"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return DB, err
}
