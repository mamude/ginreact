package main

// ListMarketsService service
func ListMarketsService() []Market {
	markets := []Market{}
	DB.Joins("CategoryBusiness").Order("Random()").Find(&markets)
	return markets
}

// GetMarketByID service
func GetMarketByID(id int) Market {
	market := Market{}
	DB.Joins("CategoryBusiness").Preload("Products").First(&market, id)
	return market
}
