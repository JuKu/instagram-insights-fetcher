package db

import (
	"database/sql"
	"fmt"

	// We need this import, but we will not use it directly
	// _ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	fmt.Println("Connecting to database ...")

	var err error
	DB, err = sql.Open("sqlite", "./data/instagram-fetcher.db") //"sqlite3"
	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createIGAccountsTable := `CREATE TABLE IF NOT EXISTS ig_accounts (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	app_id VARCHAR(255) NOT NULL,
    	account_id VARCHAR(255) NOT NULL,
    	account_name VARCHAR(255) NOT NULL,
    	api_token VARCHAR(600) NOT NULL,
    	activated BOOLEAN NOT NULL DEFAULT TRUE
	)`

	_, err := DB.Exec(createIGAccountsTable)
	if err != nil {
		panic("Could not create users table." + err.Error())
	}
}
