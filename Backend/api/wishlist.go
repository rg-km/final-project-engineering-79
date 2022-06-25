package api

import (
	"fmt"
	"net/http"
	"usedbooks/Backend/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type wishlistHandler struct {
	wishlistService repository.WRepository
}

type WishlistResponse struct {
	ID           int    `json:"id"`
	ProductID    int    `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductPrice int    `json:"price"`
	ProductImage string `json:"product_image"`
}

func NewWishlistHandler(wishlistRepository repository.WRepository) *wishlistHandler {
	return &wishlistHandler{wishlistRepository}
}

func (h *wishlistHandler) CreateWishlist(c *gin.Context) {
	var wishlistRequest repository.WishlistRequest

	if err := c.ShouldBindJSON(&wishlistRequest); err != nil {
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
	res, err := h.wishlistService.InsertWishlist(wishlistRequest)
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
