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

	router := gin.Default()
	router.POST("/register", userHandler.PostUserRegist)
	router.POST("/login", authHandler.LoginUser)
	router.GET("/users", userHandler.GetUsers)

	router.Run()
}
