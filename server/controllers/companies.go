package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mamude/ginreact/services"
	"net/http"
)

func ListCompanies(c *gin.Context) {
	_, status := services.VerifyToken(c.Request)
	if status != 200 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthoriezed!"})
		return
	}
	companies := services.ListCompanies()
	c.JSON(http.StatusOK, gin.H{"companies": companies})
}
