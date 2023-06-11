package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Connect() (*gorm.DB, error) {
	dsn := os.Getenv("DB")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
