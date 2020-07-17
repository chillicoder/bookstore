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
		connection_string = "masterbuilderusr:masterbuilderpwd8342@(masterbuilder-backend-aurora-db-cluster.cluster-cyvdcg4q9tne.ca-central-1.rds.amazonaws.com:3306)/masterbuilder_db?charset=utf8&parseTime=True"
	}

	log.Print(connection_string)
	database, err := gorm.Open("mysql", connection_string)

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
