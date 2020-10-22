package main

// ListMarketsService service
func ListMarketsService() []Market {
	var markets = []Market{}
	DB.Joins("CategoryBusiness").Find(&markets)
	return markets
}
