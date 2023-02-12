package entity

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/product?parseTime=true"))
	if err != nil {
		panic("Failed to connect database")
	}

	return database
}
