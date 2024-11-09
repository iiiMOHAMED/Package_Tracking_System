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
	DB, err = sql.Open("mysql", "root:mo1236@tcp(localhost:3306)/Ecommerce?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	// Create User table if it doesn't exist
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100) UNIQUE,
        phone VARCHAR(15),
        password VARCHAR(255),
		role ENUM('admin', 'courier', 'customer') DEFAULT 'customer' NOT NULL
    );`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	//create orders table if doesnt exist
	// Create User table if it doesn't exist
	createTableSQL2 := `CREATE TABLE IF NOT EXISTS orders (
        orderNumber INT AUTO_INCREMENT PRIMARY KEY,
        pickupLocation VARCHAR(100),
        dropOffLocation VARCHAR(100) ,
        PackageDetails VARCHAR(255),
        deliveryTime VARCHAR(100),
		user_id INT,
		courier_id INT,
		status ENUM('pending','picked up','in transit','delivered') DEFAULT 'pending' NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	    FOREIGN KEY (courier_id) REFERENCES users(id) ON DELETE SET NULL
    );`
	/*,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	    	FOREIGN KEY (courier_id) REFERENCES users(id) ON DELETE SET NULL
	*/

	_, err = DB.Exec(createTableSQL2)
	if err != nil {
		log.Fatal(err)
	}
}
