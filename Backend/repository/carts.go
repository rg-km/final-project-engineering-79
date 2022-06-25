package repository

import (
	"database/sql"
	"fmt"
)

type CRepository interface {
	FetchCart() ([]Cartlist, error)
	InsertCart(cartRequest CartRequest) (Cart, error) //untuk insert data user ke db
	DeleteCartByID(ID int) (bool, error)
}

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (c *CartRepository) FetchCart() ([]Cartlist, error) {
	sqlStatement := `SELECT c.id, c.product_id, c.quantity, p.name, p.price, p.image FROM cart_item c INNER JOIN product p ON c.product_id = p.id`

	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []Cartlist
	for rows.Next() {
		var c Cartlist
		err := rows.Scan(
			&c.ID,
			&c.ProductID,
			&c.Quantity,
			&c.ProductName,
			&c.ProductPrice,
			&c.ProductImage,
		)
		if err != nil {
			return carts, err
		}
		carts = append(carts, c)
	}
	return carts, nil
}

func (c *CartRepository) InsertCart(cartRequest CartRequest) (Cart, error) {
	sqlStatement := `INSERT INTO cart_item (product_id, quantity) VALUES (?,?)`

	res, err := c.db.Prepare(sqlStatement)

	if err != nil {
		return Cart{}, err
	}
	defer res.Close()

	newRes, err := res.Exec(
		cartRequest.ProductID,
		cartRequest.Quantity,
	)

	fmt.Println("succes", newRes)
	newCart := Cart{
		ProductID: cartRequest.ProductID,
		Quantity:  cartRequest.Quantity,
	}
	if err != nil {
		return Cart{}, err
	}

	return newCart, nil
}

func (c *CartRepository) DeleteCartByID(ID int) (bool, error) {
	// cart := Cart{}
	sqlStatement := `DELETE FROM cart_item WHERE id = ?`

	res, err := c.db.Prepare(sqlStatement)

	if err != nil {
		return false, err
	}
	defer res.Close()

	_, err = res.Exec(ID)
	if err != nil {
		return false, err
	}

	return true, nil
}
