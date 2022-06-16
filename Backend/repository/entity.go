package repository

type Product struct {
	ID    int64  `db:"id"`
	name  string `db:"name"`
	desc  string `db:"desc"`
	price int    `db:"price"`
}
