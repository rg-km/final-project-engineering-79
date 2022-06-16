package main

import (
	"database/sql"
	"usedbooks/Backend/api"
	"usedbooks/Backend/repository"
)

func main() {
	db, err := sql.Open("sqlite3", "../usedbooks.db")
	if err != nil {
		panic(err)
	}

	productsRepo := repository.NewProductRepository(db)

	mainAPI := api.NewAPI(*productsRepo)
	mainAPI.Start()
}
