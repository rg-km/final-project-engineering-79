package repository

import (
	"net/http"
	"usedbooks/backend/api"

	"github.com/gin-gonic/gin"
)

func AddUsers(c *gin.Context) {

	var json api.User

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := api.AddUser(json)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
