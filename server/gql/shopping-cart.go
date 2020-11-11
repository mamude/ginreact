package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/mamude/ginreact/services"
)

var shoppingCartType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ShoppingCart",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"market": &graphql.Field{
			Type: marketType,
		},
		"product": &graphql.Field{
			Type: productType,
		},
		"amount": &graphql.Field{
			Type: graphql.Int,
		},
		"price": &graphql.Field{
			Type: graphql.Float,
		},
		"sessionId": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var shoppingCartTotalType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ShoppingCartTotal",
	Fields: graphql.Fields{
		"amount": &graphql.Field{
			Type: graphql.Int,
		},
		"total": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

var shoppingCartArgumentFields = graphql.FieldConfigArgument{
	"customerId": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"marketId": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"productId": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"amount": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"sessionId": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

var shoppingCartListArgumentFields = graphql.FieldConfigArgument{
	"customerId": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"sessionId": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

// ShoppingCartListQuery object
var ShoppingCartListQuery = &graphql.Field{
	Type: graphql.NewList(shoppingCartType),
	Args: shoppingCartListArgumentFields,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		shoppingCart := services.ListShoppingCart(params.Args["sessionId"].(string), params.Args["customerId"].(int))
		return shoppingCart, nil
	},
}

// ShoppingCartTotalQuery object
var ShoppingCartTotalQuery = &graphql.Field{
	Type: shoppingCartTotalType,
	Args: shoppingCartListArgumentFields,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		shoppingCart := services.SumShoppingCartService(params.Args["sessionId"].(string), params.Args["customerId"].(int))
		return shoppingCart, nil
	},
}

// AddToShoppingCartMutation object
var AddToShoppingCartMutation = &graphql.Field{
	Type: shoppingCartType,
	Args: shoppingCartArgumentFields,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		shoppingCart, err := services.AddProductToShoppingCart(params)
		if err != nil {
			return nil, err
		}
		return shoppingCart, err
	},
}
