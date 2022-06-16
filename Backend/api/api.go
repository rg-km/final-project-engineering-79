package api

import (
	"fmt"
	"net/http"
	"usedbooks/Backend/repository"
)

type API struct {
	productsRepo repository.ProductRepository
	mux          *http.ServeMux
}

func NewAPI(productsRepo repository.ProductRepository) API {
	mux := http.NewServeMux()
	api := API{
		productsRepo, mux,
	}

	mux.Handle("/api/products", api.GET(api.AuthMiddleWare(http.HandlerFunc(api.productList))))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}
