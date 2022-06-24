package repository

//untuk mengakses ke db

import (
	"database/sql"
	"fmt"
)

type PRepository interface {
	FetchProduct() ([]Product, error)
	InsertProduct(productRequest ProductRequest) (Product, error) //untuk insert data user ke db
	FindProductByID(ID int) (Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) FetchProduct() ([]Product, error) {
	var products []Product

	row, err := p.db.Query("SELECT * FROM product")
	if err != nil {
		return products, err
	}
	for row.Next() {
		var product Product

		err := row.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Image,
		)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductRepository) InsertProduct(productRequest ProductRequest) (Product, error) {
	sqlStatement := `INSERT INTO product (name, description, price, image) VALUES (?,?,?,?)`

	res, err := p.db.Prepare(sqlStatement)

	if err != nil {
		return Product{}, err
	}
	defer res.Close()

	newRes, err := res.Exec(
		productRequest.Name,
		productRequest.Description,
		productRequest.Price,
		productRequest.Image,
	)
	fmt.Println("succes", newRes)
	newProduct := Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		Image:       productRequest.Image,
	}
	if err != nil {
		return Product{}, err
	}

	return newProduct, nil
}

func (p *ProductRepository) FindProductByID(ID int) (Product, error) {
	product := Product{}
	sqlStatement := `SELECT * FROM product WHERE id = ?`

	row := p.db.QueryRow(sqlStatement, ID)
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Image,
	)
	if err != nil {
		return product, err
	}
	return product, nil
}
