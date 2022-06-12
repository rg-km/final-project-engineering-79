package api

import (
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AddUser(newUser User) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO users (nama, email, password) VALUES (?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newUser.Nama, newUser.Email, newUser.Password)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
