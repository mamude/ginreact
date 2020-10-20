package main

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/graphql-go/graphql"
)

type personDto struct {
	ID        int     `json:"id"`
	Age       int     `json:"age" validate:"required,gte=18"`
	Level     string  `json:"level" validate:"required"`
	Country   string  `json:"country" validate:"required"`
	FirstName string  `json:"firstName" validate:"required"`
	LastName  string  `json:"lastName" validate:"required"`
	Salary    float64 `json:"salary" validate:"required"`
}

var personType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Person",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
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

var personArgumentFields = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"age": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"level": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"country": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"firstName": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"lastName": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"salary": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Float),
	},
}

// PersonQuery object
var PersonQuery = &graphql.Field{
	Type: graphql.NewList(personType),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		people := ListPeopleService()
		return people, nil
	},
}

// PersonCreateMutation object
var PersonCreateMutation = &graphql.Field{
	Type: personType,
	Args: personArgumentFields,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		age, _ := params.Args["age"].(int)
		level, _ := params.Args["level"].(string)
		country, _ := params.Args["country"].(string)
		firstName, _ := params.Args["firstName"].(string)
		lastName, _ := params.Args["lastName"].(string)
		salary, _ := params.Args["salary"].(float64)

		v := validator.New()

		personDto := personDto{
			Age:       age,
			Level:     level,
			Country:   country,
			FirstName: firstName,
			LastName:  lastName,
			Salary:    salary,
		}

		if err := v.Struct(personDto); err != nil {
			return nil, errors.New(err.Error())
		}

		person := Person{
			Age:       personDto.Age,
			Level:     personDto.Level,
			Country:   personDto.Country,
			FirstName: personDto.FirstName,
			LastName:  personDto.LastName,
			Salary:    personDto.Salary,
		}

		if err := DB.Create(&person).Error; err != nil {
			return nil, errors.New(err.Error())
		}

		return person, nil
	},
}

// PersonUpdateMutation object
var PersonUpdateMutation = &graphql.Field{
	Type: personType,
	Args: personArgumentFields,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := params.Args["id"].(int)
		age, _ := params.Args["age"].(int)
		level, _ := params.Args["level"].(string)
		country, _ := params.Args["country"].(string)
		firstName, _ := params.Args["firstName"].(string)
		lastName, _ := params.Args["lastName"].(string)
		salary, _ := params.Args["salary"].(float64)

		v := validator.New()

		personDto := personDto{
			ID:        id,
			Age:       age,
			Level:     level,
			Country:   country,
			FirstName: firstName,
			LastName:  lastName,
			Salary:    salary,
		}

		if err := v.Struct(personDto); err != nil {
			return nil, errors.New(err.Error())
		}

		person := Person{}
		DB.First(&person, id)
		person.Age = personDto.Age
		person.Level = personDto.Level
		person.Country = personDto.Country
		person.FirstName = personDto.FirstName
		person.LastName = personDto.LastName
		person.Salary = personDto.Salary

		if err := DB.Save(&person).Error; err != nil {
			return nil, errors.New(err.Error())
		}

		return person, nil
	},
}

// PersonDeleteMutation object
var PersonDeleteMutation = &graphql.Field{
	Type: personType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := params.Args["id"].(int)

		if err := DB.Delete(&Person{}, id).Error; err != nil {
			return nil, errors.New(err.Error())
		}

		return nil, nil
	},
}
