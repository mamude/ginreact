package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/mamude/ginreact/services"
)

var categoryBusinessType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CategoryBusiness",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var marketType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Market",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"url": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"phone": &graphql.Field{
			Type: graphql.String,
		},
		"rating": &graphql.Field{
			Type: graphql.Float,
		},
		"distance": &graphql.Field{
			Type: graphql.Float,
		},
		"deliveryTime": &graphql.Field{
			Type: graphql.String,
		},
		"deliveryTax": &graphql.Field{
			Type: graphql.Float,
		},
		"categoryBusiness": &graphql.Field{
			Type: categoryBusinessType,
		},
		"products": &graphql.Field{
			Type: graphql.NewList(productType),
		},
	},
})

var marketArgumentFields = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
}

// MarketListQuery object
var MarketListQuery = &graphql.Field{
	Type: graphql.NewList(marketType),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		markets := services.ListMarketsService()
		return markets, nil
	},
}

// MarketByIDQuery object
var MarketByIDQuery = &graphql.Field{
	Type: marketType,
	Args: marketArgumentFields,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		market := services.GetMarketByIDService(params.Args["id"].(int))
		return market, nil
	},
}
