package main

import (
	"log"
	"usedbooks/backend/api"
	"usedbooks/backend/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	err := api.ConnectDatabase()
	checkErr(err)
	r := gin.Default()

	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.POST("users", repository.AddUsers)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
