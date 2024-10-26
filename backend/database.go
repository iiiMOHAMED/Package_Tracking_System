package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Change username, password, and dbname as needed
	DB, err = sql.Open("mysql", "root:ahmed342002@tcp(localhost:3306)/Ecommerce?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	// Create User table if it doesn't exist
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100) UNIQUE,
        phone VARCHAR(15),
        password VARCHAR(255)
    );`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
