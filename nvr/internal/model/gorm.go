package model

import (
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() {
	db, err := gorm.Open(sqlite.Open("nvr.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("[DB ERROR] : ", err)
	}
	err = db.AutoMigrate(&Role{}, &User{})
	if err != nil {
		log.Fatalln("[DB ERROR] : ", err)
	}
	DB = db
}
