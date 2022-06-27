package api

import (
	"fmt"
	"net/http"
	"strconv"
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

func (h *wishlistHandler) GetWishlists(c *gin.Context) {
	wishlists, err := h.wishlistService.FetchWishlist()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	var wishlistsResponse []WishlistResponse

	for _, v := range wishlists {
		wishlistResponse := convertToWishlistResponse(v)

		wishlistsResponse = append(wishlistsResponse, wishlistResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": wishlistsResponse,
	})
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

func (h *wishlistHandler) DeleteWishlist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Invalid ID",
		})
	}
	res, err := h.wishlistService.DeleteWishlistByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "Wishlist deleted successfully",
	})
}

func (h *wishlistHandler) UpdateWishlist(c *gin.Context) {
	var json repository.Wishlist

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Invalid ID",
		})
	}

	success, err := h.wishlistService.UpdatedWishlist(json, id)

	if success {
		c.JSON(http.StatusOK, gin.H{
			"message": "Sucess Updated",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errors": err,
		})
	}
}

func convertToWishlistResponse(h repository.Wishlist) WishlistResponse {
	return WishlistResponse{
		ID:           h.ID,
		ProductID:    h.ProductID,
		ProductName:  h.ProductName,
		ProductPrice: h.ProductPrice,
		ProductImage: h.ProductImage,
	}
}
