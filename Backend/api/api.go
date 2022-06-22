package api

import (
	"fmt"
	"net/http"
	"usedbooks/Backend/repository"
)

type API struct {
	usersRepo    repository.UserRepository
	productsRepo repository.ProductRepository
	mux          *http.ServeMux
}

func NewAPI(usersRepo repository.UserRepository, productsRepo repository.ProductRepository) API {
	mux := http.NewServeMux()
	api := API{}

	// Handler untuk users (dengan middleware)
	// ...

	// Handler untuk guest (tanpa middleware)
	mux.Handle("/api/product/product_name", api.GET(http.HandlerFunc(api.getProductByName))) //still red
	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}
