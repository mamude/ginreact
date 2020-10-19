package main

import (
	"github.com/graphql-go/graphql"
)

var companyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Company",
	Fields: graphql.Fields{
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

// CompanyQuery query object
var CompanyQuery = &graphql.Field{
	Type:        graphql.NewList(companyType),
	Description: "List Companies",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		companies := ListCompaniesService()
		return companies, nil
	},
}
