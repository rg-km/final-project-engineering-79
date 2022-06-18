package api

type Product struct {
	ID          int`json:"id"`
	ProductName string`json:"product_name"`
	Desc        string`json:"desc"`
	Price       int	`json:"price"`
}

type ProductListSuccessResponse struct {
	Product []Product `json:"product"`
}

// func (api *API) productList(w http.ResponseWriter, req *http.Request) {

// }
