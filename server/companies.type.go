package main

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/graphql-go/graphql"
)

type companyDto struct {
	ID          int    `json:"id"`
	OpeningYear int    `json:"openingYear" validate:"required,gte=1930"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

var companyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Company",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
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

// CompanyCreateMutation object
var CompanyCreateMutation = &graphql.Field{
	Type: companyType,
	Args: graphql.FieldConfigArgument{
		"openingYear": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		openingYear, _ := params.Args["openingYear"].(int)
		name, _ := params.Args["name"].(string)
		description, _ := params.Args["description"].(string)

		v := validator.New()

		companyDto := companyDto{
			Name:        name,
			Description: description,
			OpeningYear: openingYear,
		}

		if err := v.Struct(companyDto); err != nil {
			return nil, errors.New(err.Error())
		}

		company := Company{
			Name:        &companyDto.Name,
			Description: companyDto.Description,
			OpeningYear: companyDto.OpeningYear,
		}

		if err := DB.Create(&company).Error; err != nil {
			return nil, errors.New("This company already exists")
		}

		return company, nil
	},
}
