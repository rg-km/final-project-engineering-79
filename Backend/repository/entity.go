package repository

// sama dengan field pada db
//untuk menampung data pada golang butuh sebuah struct

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
