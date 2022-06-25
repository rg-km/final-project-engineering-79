package repository

import "time"

// sama dengan field pada db
//untuk menampung data pada golang butuh sebuah struct

type User struct {
	ID        int
	Username  string
	Email     string
	Nohp      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
	Image       string
}

type Cart struct {
	ID        int
	ProductID int
	Quantity  int
}

type Cartlist struct {
	ID           int
	ProductID    int
	ProductName  string
	ProductPrice int
	ProductImage string
	Quantity     int
}

type Wishlist struct {
	ID        int
	ProductID int
}
