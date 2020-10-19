package main

import (
	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"companies": CompanyQuery,
		"people":    PersonQuery,
	},
})

// CompanySchema Graphql Schema
var CompanySchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
