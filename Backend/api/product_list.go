package api

import (
	"encoding/json"
	"net/http"
)

type ProductErrorResponse struct {
	Error string `json:"error"`
}

type Product struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	Desc        string `json:"desc"`
	Price       int    `json:"price"`
}

type ProductSuccessResponse struct {
	Products []Product `json:"products"`
}

func (api *API) getProductCategory(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	product_name, err := api.productsRepo.FetchProductByName() // still red
	encoder := json.NewEncoder(w)
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(ProductErrorResponse{Error: err.Error()})
		}
	}()

	encoder.Encode(product_name)
}
