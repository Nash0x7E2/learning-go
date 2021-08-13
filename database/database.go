package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() *gorm.DB {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

	if err != nil {
		panic("Unable to connect to database")
	}

	fmt.Println("Database connection open 🚀")
	return DBConn
}
