package main

import (
	"database/sql"
	"usedbooks/Backend/api"
	"usedbooks/Backend/repository"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db/usedbooks.db")
	if err != nil {
		panic(err)
	}

	//dependency
	userRepository := repository.NewUserRepository(db)
	userHandler := api.NewUserHandler(userRepository)
	authHandler := api.NewAuthHandler(userRepository)

	productRepository := repository.NewProductRepository(db)
	productHandler := api.NewProductHandler(productRepository)

	cartRepository := repository.NewCartRepository(db)
	cartHandler := api.NewCartHandler(cartRepository)

	router := gin.Default()
	router.POST("/register", userHandler.PostUserRegist)
	router.POST("/login", authHandler.LoginUser)
	router.GET("/users", userHandler.GetUsers)

	//product
	router.POST("/insertProduct", productHandler.CreateProduct)
	router.GET("/listProduct", productHandler.GetProducts)
	router.GET("/detailProduct/:id", productHandler.GetProduct)

	//cart
	router.POST("/insertCart", cartHandler.CreateCart)
	router.GET("/listCarts", cartHandler.GetCarts)
	router.DELETE("/deleteCart/:id", cartHandler.DeleteCart)

	router.Run()
}
