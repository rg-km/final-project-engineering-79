package api

import (
	"encoding/json"
	"net/http"
)

type ProductListErrorRespones struct {
	Error string `json:"error"`
}

type Product struct {
	ID    int64  `json:"id"`
	name  string `json:"name"`
	desc  string `json:"desc"`
	price int    `json:"price"`
	// image
}

type ProductListSuccessRespones struct {
	Products []Product `json:"products"`
}

func (api *API) productList(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)

	response := ProductListSuccessRespones{}
	response.Products = make([]Product, 0)

	products, err := api.productsRepo.FetchProducts()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(DashboardErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	for _, product := range products {
		response.Products = append(response.Products, Product{
			ID:    product.ID,
			name:  product.name,
			desc:  product.desc,
			price: product.price,
		})
	}
	encoder.Encode(response)
}
