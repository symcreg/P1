package db

import (
	"P1/config"
	"github.com/jinzhu/gorm"
)

func InitDB() {
	db, err := gorm.Open("sqlite3", "P1.db")
	if err != nil {
		panic("connect database error")
	}
	defer db.Close()
	db.AutoMigrate(&config.User{})
}
