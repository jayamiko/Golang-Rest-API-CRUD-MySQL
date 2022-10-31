package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:qwerty123@tcp(127.0.0.1:3306)/rest_api_go"))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection Database Success")
	}
	db.AutoMigrate(&Product{})
	DB = db
}
