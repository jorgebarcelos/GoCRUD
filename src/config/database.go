package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err :=  sql.Open("mysql", "sql5677686:WQlnHxs4ZR@tcp(sql5.freesqldatabase.com)/sql5677686?parseTime=true")

	if err != nil {
		panic(err)
	}

	log.Println("Database connected")
	DB = db
}