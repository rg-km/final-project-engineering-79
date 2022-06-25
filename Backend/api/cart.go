package api

import (
	"fmt"
	"net/http"
	"strconv"
	"usedbooks/Backend/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type cartHandler struct {
	cartService repository.CRepository
}

type CartResponse struct {
	ID           int    `json:"id"`
	ProductID    int    `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductPrice int    `json:"price"`
	ProductImage string `json:"product_image"`
	Quantity     int    `json:"quantity"`
}

func NewCartHandler(cartRepository repository.CRepository) *cartHandler {
	return &cartHandler{cartRepository}
}

func (h *cartHandler) GetCarts(c *gin.Context) {
	carts, err := h.cartService.FetchCart()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	var cartsResponse []CartResponse

	for _, v := range carts {
		cartResponse := convertToCartResponse(v)

		cartsResponse = append(cartsResponse, cartResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": cartsResponse,
	})
}

func (h *cartHandler) CreateCart(c *gin.Context) {
	var cartRequest repository.CartRequest

	if err := c.ShouldBindJSON(&cartRequest); err != nil {
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
	res, err := h.cartService.InsertCart(cartRequest)
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

func (h *cartHandler) DeleteCart(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Invalid ID",
		})
	}
	res, err := h.cartService.DeleteCartByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "Delete Cart Success",
	})
}

func (h *cartHandler) UpdateCart(c *gin.Context) {
	var json repository.Cart

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
	}

	success, err := h.cartService.UpdatedCart(json, id)

	if success {
		c.JSON(http.StatusOK, gin.H{
			"message": "Success Updated",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
	}
}

func convertToCartResponse(h repository.Cartlist) CartResponse {
	return CartResponse{
		ID:           h.ID,
		ProductID:    h.ProductID,
		ProductName:  h.ProductName,
		ProductPrice: h.ProductPrice,
		ProductImage: h.ProductImage,
		Quantity:     h.Quantity,
	}
}
