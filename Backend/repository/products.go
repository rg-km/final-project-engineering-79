package repository

import "database/sql"

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) FetchProductByID(id int) (Product, error) {
	var product Product
	var sqlStatement string

	sqlStatement = `
		SELECT id, product_name, desc, price
		FROM product WHERE id =?
	`

	err := p.db.QueryRow(sqlStatement, id).Scan(
		&product.ID,
		&product.ProductName,
		&product.Desc,
		&product.Price,
	)

	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *ProductRepository) FetchProductByName(ProductName string) (Product, error) {
	var product Product
	var sqlStatement string

	sqlStatement = `
		SELECT id, product_name, desc, price
		FROM product WHERE name =?
	`

	err := p.db.QueryRow(sqlStatement, ProductName).Scan(
		&product.ID,
		&product.ProductName,
		&product.Desc,
		&product.Price,
	)

	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *ProductRepository) FetchProducts() ([]Product, error) {
	products := []Product{}

	rows, err := p.db.Query("SELECT id, product_name, desc, price FROM products")

	defer rows.Close()

	for rows.Next() {
		var product Product
		if err := rows.Scan(
			&product.ID,
			&product.ProductName,
			&product.Desc,
			&product.Price,
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
