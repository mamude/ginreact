package services

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/mamude/ginreact/models"
)

// SyncShoppingCart service
func SyncShoppingCart(sessionID string, customerID int) {
	shoppingCart := []models.ShoppingCart{}
	models.DB.Where("session_id = ?", sessionID).Find(&shoppingCart)
	for _, cart := range shoppingCart {
		cart.CustomerID = customerID
		cart.SessionID = ""
		models.DB.Save(&cart)
	}
}

// ListShoppingCart service
func ListShoppingCart(sessionID string, customerID int) []models.ShoppingCart {
	shoppingCart := []models.ShoppingCart{}
	if sessionID != "" {
		models.DB.Joins("Market").Joins("Product").Where("session_id = ?", sessionID).Find(&shoppingCart)
		return shoppingCart
	}
	models.DB.Joins("Market").Joins("Product").Where("customer_id = ?", customerID).Find(&shoppingCart)
	return shoppingCart
}

// AddProductToShoppingCart service
func AddProductToShoppingCart(params graphql.ResolveParams) (models.ShoppingCart, error) {
	shoppingCart := models.ShoppingCart{}
	product := models.Product{}

	productID := params.Args["productId"].(int)
	customerID := params.Args["customerId"].(int)
	marketID := params.Args["marketId"].(int)
	amount := params.Args["amount"].(int)
	sessionID := params.Args["sessionId"].(string)

	// check if product is available
	if err := models.DB.First(&product, productID).Error; err != nil {
		return shoppingCart, errors.New("Produto não encontrado")
	}
	if product.Amount < amount {
		return shoppingCart, errors.New("Quantidade insuficiente em estoque")
	}

	// add product to shopping cart
	shoppingCart.CustomerID = customerID
	shoppingCart.MarketID = marketID
	shoppingCart.ProductID = productID
	shoppingCart.Price = product.Price
	shoppingCart.Amount = amount
	shoppingCart.SessionID = sessionID
	models.DB.Save(&shoppingCart)

	return shoppingCart, nil
}

// SumShoppingCartService service
func SumShoppingCartService(sessionID string, customerID int) map[string]interface{} {
	result := map[string]interface{}{}
	query := "SUM(amount) as amount, SUM(price * amount) as total"
	if sessionID != "" {
		models.DB.Select(query).Model(models.ShoppingCart{}).Where("session_id = ?", sessionID).Find(&result)
		return result
	}
	models.DB.Select(query).Model(models.ShoppingCart{}).Where("customer_id = ?", customerID).Find(&result)
	return result
}
