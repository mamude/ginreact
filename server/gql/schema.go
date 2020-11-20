package gql

import (
	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"market":          MarketByIDQuery,
		"markets":         MarketListQuery,
		"shoppingCart":    ShoppingCartListQuery,
		"shoppingCartSum": ShoppingCartTotalQuery,
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"addToCart":      AddToShoppingCartMutation,
		"updateItemCart": UpdateProductShoppingCartMutation,
		"removeItemCart": RemoveProductShoppingCartMutation,
	},
})

// EcommerceSchema Graphql Schema
var EcommerceSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
