package main

import (
	"github.com/graphql-go/graphql"
)

var companyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Company",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"openingYear": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// CompanyQuery object
var CompanyQuery = &graphql.Field{
	Type: graphql.NewList(companyType),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		companies := ListCompaniesService()
		return companies, nil
	},
}
