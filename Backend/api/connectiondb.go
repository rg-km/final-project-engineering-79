package api

import "database/sql"

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./db/usedbooks.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}
