package api

import (
	"fmt"
	"net/http"
	"strconv"
	"usedbooks/Backend/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type productHandler struct {
	productService repository.PRepository
}

type ProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
}

func NewProductHandler(productRepository repository.PRepository) *productHandler {
	return &productHandler{productRepository}
}

func (h *productHandler) GetProducts(c *gin.Context) {
	products, err := h.productService.FetchProduct()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	var productsResponse []ProductResponse

	for _, v := range products {
		productResponse := convertToProductResponse(v)

		productsResponse = append(productsResponse, productResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": productsResponse,
	})
}

func (h *productHandler) GetProduct(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	product, err := h.productService.FindProductByID(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var productRequest repository.ProductRequest

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	res, err := h.productService.InsertProduct(productRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func convertToProductResponse(h repository.Product) ProductResponse {
	return ProductResponse{
		ID:          h.ID,
		Name:        h.Name,
		Description: h.Description,
		Price:       h.Price,
		Image:       h.Image,
	}
}
