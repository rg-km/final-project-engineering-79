package repository

//input dari user

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Nohp     string `json:"nohp" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Image       string `json:"image" binding:"required"`
}

type CartRequest struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

type WishlistRequest struct {
	ProductID int `json:"product_id" binding:"required"`
}
