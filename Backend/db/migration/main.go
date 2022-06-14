package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Run This Script for migration db
func main() {
	db, err := sql.Open("sqlite3", "../usedbooks.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
    id integer not null primary key AUTOINCREMENT,
    username varchar(255) not null,
    password varchar(255) not null
);

INSERT INTO users (username, password) VALUES
    ('aditira', '1234'),
    ('dina', '4321'),
    ('dito', '2552');`)

	if err != nil {
		panic(err)
	}
	defer db.Close()
}
