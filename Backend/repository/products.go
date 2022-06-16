package repository

import (
	"database/sql"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) FetchProductByID(id int64) (Product, error) {
	var product Product
	var sqlStatement string

	sqlStatement = `
		SELECT id, name, desc, price
		FROM products WHERE id = ?
	`

	err := p.db.QueryRow(sqlStatement, id).Scan(
		&product.ID,
		&product.name,
		&product.desc,
		&product.price,
	)

	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *ProductRepository) FetchProductByName(productName string) (Product, error) {
	var product Product
	var sqlStatement string

	sqlStatement = `
		SELECT id, product_name, price
		FROM products WHERE product_name = ?
	`

	err := p.db.QueryRow(sqlStatement, productName).Scan(
		&product.ID,
		&product.name,
		&product.desc,
		&product.price,
	)

	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *ProductRepository) FetchProducts() ([]Product, error) {
	products := []Product{}

	rows, err := p.db.Query("SELECT id, product_name, price FROM products")

	defer rows.Close()

	for rows.Next() {
		var product Product
		if err := rows.Scan(
			&product.ID,
			&product.name,
			&product.desc,
			&product.price,
		); err != nil {
			return products, err
		}
		products = append(products, product)
	}
	if err != nil {
		return products, err
	}
	return products, nil
}
