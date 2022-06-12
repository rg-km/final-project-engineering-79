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
    nama varchar(255) not null,
	email varchar(255) not null,
    password varchar(255) not null
);

INSERT INTO users(nama, email, password) VALUES
    ('aditira', 'adit@gmail.com', 'adit@'),
    ('dina', 'dina@gmail.com', 'employee'),
    ('dito', 'dito@gmail.com', 'employee');`)

	if err != nil {
		panic(err)
	}
}
