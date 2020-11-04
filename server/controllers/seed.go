package controllers

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/gin-gonic/gin"
	"github.com/mamude/ginreact/models"
	"github.com/mamude/ginreact/security"
	"strings"
)

func resetTables() {
	models.DB.Exec("DELETE FROM users;")
	models.DB.Exec("DELETE FROM people;")
	models.DB.Exec("DELETE FROM companies;")
	models.DB.Exec("DELETE FROM roles;")
	models.DB.Exec("DELETE FROM order_items;")
	models.DB.Exec("DELETE FROM orders;")
	models.DB.Exec("DELETE FROM shopping_carts;")
	models.DB.Exec("DELETE FROM products;")
	models.DB.Exec("DELETE FROM markets;")
	models.DB.Exec("DELETE FROM categories;")
	models.DB.Exec("DELETE FROM category_businesses;")
	models.DB.Exec("DELETE FROM customers;")
}

func seedTables() {
	// create User
	password := "admin"
	hash, _ := security.HashPassword(password)
	models.DB.Create(&models.User{Username: "admin", Password: hash})

	// create category business
	categoryBusiness := []models.CategoryBusiness{
		{Name: "Category 1"},
		{Name: "Category 2"},
		{Name: "Category 3"},
		{Name: "Category 4"},
		{Name: "Category 5"},
		{Name: "Category 6"},
		{Name: "Category 7"},
		{Name: "Category 8"},
		{Name: "Category 9"},
		{Name: "Category 10"},
	}
	models.DB.Create(&categoryBusiness)

	// create markets
	for i := 0; i < 20; i++ {
		appName := gofakeit.AppName()
		categoryBusiness := models.CategoryBusiness{}
		models.DB.First(&categoryBusiness, gofakeit.Number(1, 10))
		models.DB.Create(&models.Market{
			CategoryBusinessID: categoryBusiness.ID,
			Name:               &appName,
			Description:        gofakeit.Phrase(),
			Email:              gofakeit.Email(),
			URL:                gofakeit.URL(),
			Phone:              gofakeit.Phone(),
			Rating:             gofakeit.Float64Range(1, 5),
			Distance:           gofakeit.Float64Range(1, 6),
			DeliveryTime:       "30-40min",
			DeliveryTax:        gofakeit.Price(1, 6),
		})
	}

	// create categories
	categories := []models.Category{
		{Name: "Category 1"},
		{Name: "Category 2"},
		{Name: "Category 3"},
		{Name: "Category 4"},
		{Name: "Category 5"},
		{Name: "Category 6"},
		{Name: "Category 7"},
		{Name: "Category 8"},
		{Name: "Category 9"},
		{Name: "Category 10"},
	}
	models.DB.Create(&categories)

	// create products
	for i := 0; i < 100; i++ {
		market := models.Market{}
		category := models.Category{}
		models.DB.First(&category, gofakeit.Number(1, 10))
		models.DB.First(&market, gofakeit.Number(1, 10))
		productName := []string{
			gofakeit.BeerName(),
			gofakeit.BeerMalt(),
			gofakeit.Breakfast(),
			gofakeit.Fruit(),
			gofakeit.Vegetable(),
			gofakeit.Lunch(),
			gofakeit.Snack(),
		}
		models.DB.Create(&models.Product{
			CategoryID:  category.ID,
			MarketID:    market.ID,
			Name:        gofakeit.RandomString(productName),
			Description: gofakeit.LoremIpsumSentence(50),
			Price:       gofakeit.Price(1, 50),
			Amount:      gofakeit.Number(1, 100),
		})
	}

	// create customers
	for i := 0; i < 100; i++ {
		lastName := gofakeit.LastName()
		email := strings.ToLower(lastName + "@" + gofakeit.DomainName())
		password, _ = security.HashPassword("123456")
		models.DB.Create(&models.Customer{
			FirstName: gofakeit.FirstName(),
			LastName:  lastName,
			Phone:     gofakeit.Phone(),
			Email:     &email,
			Password:  password,
		})
	}
}

// ResetDatabaseHandler action
func ResetDatabaseHandler(c *gin.Context) {
	go resetTables()
	c.JSON(200, gin.H{"message": "Reseting database..."})
}

// SeedDataHandler action
func SeedDataHandler(c *gin.Context) {
	go seedTables()
	c.JSON(200, gin.H{"message": "Adding fake data..."})
}
