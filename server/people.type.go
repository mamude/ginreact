package main

import (
	"github.com/graphql-go/graphql"
)

var personType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Person",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"age": &graphql.Field{
			Type: graphql.Int,
		},
		"level": &graphql.Field{
			Type: graphql.String,
		},
		"country": &graphql.Field{
			Type: graphql.String,
		},
		"firstName": &graphql.Field{
			Type: graphql.String,
		},
		"lastName": &graphql.Field{
			Type: graphql.String,
		},
		"salary": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

// PersonQuery object
var PersonQuery = &graphql.Field{
	Type: graphql.NewList(personType),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		people := ListPeopleService()
		return people, nil
	},
}
