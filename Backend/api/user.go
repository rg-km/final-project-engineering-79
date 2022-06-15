package api

import "database/sql"

func searchUser(db *sql.DB, Id int) user {
	rows, _ := db.Query("SELECT * FROM users WHERE Id = ?", Id)
	defer rows.Close()

}
