package main

import (
	"database/sql"
	"usedbooks/Backend/repository"
	"usedbooks/backend/api"

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

	router := gin.Default()
	router.POST("/register", userHandler.PostUserRegist)
	router.POST("/login", authHandler.LoginUser)
	router.GET("/users", userHandler.GetUsers)

	//product
	router.POST("/insertProduct", productHandler.CreateProduct)
	router.GET("/listProduct", productHandler.GetProducts)

	router.Run()
}
