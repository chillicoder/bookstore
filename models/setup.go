// models/setup.go

package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	var connection_string, exists = os.LookupEnv("DB_URL")

	if !exists {
		connection_string = "bookstore:bookstorepwd1@localhost:3306/bookstore_db?charset=utf8&parseTime=True"
	}

	log.Print(connection_string)
	database, err := gorm.Open("mysql", connection_string)

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
