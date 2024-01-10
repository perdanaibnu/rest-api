package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/belajar_golang_db"))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Item{})

	DB = db
}
