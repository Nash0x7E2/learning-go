package database

import (
	"fmt"
	"learning_go/book"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

	if err != nil {
		panic("Unable to connect to database")
		os.Exit(2)
	}

	fmt.Println("Database connection open 🚀")
	db.Logger = logger.Default.LogMode(logger.Info)

	fmt.Println("Running migrations")
	db.AutoMigrate(&book.Book{})

	DB = DBInstance{
		Db: db,
	}
}
