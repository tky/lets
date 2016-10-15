package db

import (
	"lets/models"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *gorm.DB {

	DB, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Product{})

	return DB
}
