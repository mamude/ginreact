package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mamude/ginreact/services"
	"net/http"
)

func ListCompanies(c *gin.Context) {
	companies := services.ListCompanies()
	c.JSON(http.StatusOK, gin.H{"companies": companies})
}
