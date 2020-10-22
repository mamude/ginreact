package main

import (
	"github.com/graphql-go/graphql"
)

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
	},
})

// MarketListQuery object
var MarketListQuery = &graphql.Field{
	Type: graphql.NewList(marketType),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		markets := ListMarketsService()
		return markets, nil
	},
}
