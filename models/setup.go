package models

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// var DB *gorm.DB

func ConnectDatabase() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:qwerty123@tcp(127.0.0.1:3306)/rest_api_go")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connection Database Success")
	}
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(60 * time.Second)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	return db
}

var DB = ConnectDatabase()