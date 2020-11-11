package gql

import (
	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"market":  MarketByIDQuery,
		"markets": MarketListQuery,
	},
})

// var rootMutation = graphql.NewObject(graphql.ObjectConfig{
// 	Name:   "RootMutation",
// 	Fields: graphql.Fields{},
// })

// EcommerceSchema Graphql Schema
var EcommerceSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
	//Mutation: rootMutation,
})
