package services

import (
	"github.com/mamude/ginreact/models"
)

// ListMarketsService service
func ListMarketsService() []models.Market {
	markets := []models.Market{}
	models.DB.Joins("CategoryBusiness").Order("Random()").Find(&markets)
	return markets
}

// GetMarketByIDService service
func GetMarketByIDService(id int) models.Market {
	market := models.Market{}
	models.DB.Joins("CategoryBusiness").Preload("Products").First(&market, id)
	return market
}
