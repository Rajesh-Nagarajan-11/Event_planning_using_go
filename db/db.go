package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")
	if err != nil {
		panic(err)
	}

	// Optional DB settings
	DB.SetConnMaxLifetime(time.Hour)
	DB.SetMaxIdleConns(5)

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to Database Successfully ")
	CreateTable()
}

func CreateTable() {
	createUserTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic(err) // print the real error
	}
	log.Println("Users Table Created/Verified")

	createEventTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		userid INTEGER,
		FOREIGN KEY (userid) REFERENCES users(id)
	);`

	_, err = DB.Exec(createEventTable)
	if err != nil {
		panic(err) // print the real error
	}
	log.Println("Events Table Created/Verified")
}
