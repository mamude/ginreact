package main

import (
	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"market":    MarketByIDQuery,
		"markets":   MarketListQuery,
		"companies": CompanyQuery,
		"people":    PersonQuery,
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		// company mutations
		"createCompany": CompanyCreateMutation,
		"updateCompany": CompanyUpdateMutation,
		"deleteCompany": CompanyDeleteMutation,
		// person mutations
		"createPerson": PersonCreateMutation,
		"updatePerson": PersonUpdateMutation,
		"deletePerson": PersonDeleteMutation,
	},
})

// EcommerceSchema Graphql Schema
var EcommerceSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
