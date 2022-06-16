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
	CREATE TABLE IF NOT EXISTS products (
    id integer not null primary key AUTOINCREMENT,
    name varchar(255) not null,
	desc varchar(255) not null,
    price integer(32) not null
);

INSERT INTO users(nama, email, password) VALUES
    ('Crazy Love', 'Ini adalah novel', '170000'),
    ('Your Heart Is The Sea', 'Ini adalah novel', '120000'),
    ('Milk and Honey', 'Ini adalah novel susu dan madu', '160000');`)

	if err != nil {
		panic(err)
	}
}
