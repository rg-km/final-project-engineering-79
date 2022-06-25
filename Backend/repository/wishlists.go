package repository

import (
	"database/sql"
	"fmt"
)

type WRepository interface {
	FetchWishlist() ([]Wishlist, error)
	InsertWishlist(wishlistRequest WishlistRequest) (Wishlist, error) //untuk insert data user ke db
}

type WishlistRepository struct {
	db *sql.DB
}

func NewWishlistRepository(db *sql.DB) *WishlistRepository {
	return &WishlistRepository{db: db}
}

func (w *WishlistRepository) FetchWishlist() ([]Wishlist, error) {
	sqlStatement := `SELECT w.id, w.product_id, p.name, p.price, p.image FROM wishlist_item w INNER JOIN product p ON w.product_id = p.id`

	rows, err := w.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wishlists []Wishlist
	for rows.Next() {
		var w Wishlist
		err := rows.Scan(
			&w.ID,
			&w.ProductID,
			&w.ProductName,
			&w.ProductPrice,
			&w.ProductImage,
		)
		if err != nil {
			return wishlists, err
		}
		wishlists = append(wishlists, w)
	}
	return wishlists, nil
}

func (w *WishlistRepository) InsertWishlist(wishlistRequest WishlistRequest) (Wishlist, error) {
	sqlStatement := `INSERT INTO wishlist_item (product_id) VALUES (?)`

	res, err := w.db.Prepare(sqlStatement)

	if err != nil {
		return Wishlist{}, err
	}
	defer res.Close()

	newRes, err := res.Exec(
		wishlistRequest.ProductID,
	)

	fmt.Println("success", newRes)
	newWishlist := Wishlist{
		ProductID: wishlistRequest.ProductID,
	}
	if err != nil {
		return Wishlist{}, err
	}
	return newWishlist, nil
}
