package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "main.db")
	if err != nil {
		log.Println("Error on connect to db.")
		panic(err)
	}
	db.AutoMigrate(&User{}, &Type{}, &CardSet{}, &Card{})
}
