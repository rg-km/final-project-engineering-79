package api

import (
	"github.com/gin-gonic/gin"
)

type wishlistHandler struct {
}

type WishlistResponse struct {
	ID           int    `json:"id"`
	ProductID    int    `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductPrice int    `json:"price"`
	ProductImage string `json:"product_image"`
}

func AddWishlist(c *gin.Context) {

}
