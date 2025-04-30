package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	fmt.Println("Database connection established")

	return db, nil
}
