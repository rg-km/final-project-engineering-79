package repository

//untuk mengakses ke db

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	FetchUsers() ([]User, error)
	InsertUser(userRequest UserRequest) (User, error) //untuk insert data user ke db
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUsers() ([]User, error) {
	var users []User

	row, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}
	for row.Next() {
		var user User

		err := row.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
		)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepository) InsertUser(userRequest UserRequest) (User, error) {
	sqlStatement := `INSERT INTO users (username, password) VALUES (?,?)`

	res, err := u.db.Prepare(sqlStatement)

	if err != nil {
		return User{}, err
	}
	defer res.Close()

	newRes, err := res.Exec(
		userRequest.Username,
		userRequest.Password,
	)
	fmt.Println("succes", newRes)
	newUser := User{
		Username: userRequest.Username,
		Password: userRequest.Password,
	}
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}
