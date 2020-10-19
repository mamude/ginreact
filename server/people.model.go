package main

import "gorm.io/gorm"

// Person model
type Person struct {
	gorm.Model
	Age       int
	Level     string
	Country   string
	FirstName string
	LastName  string
	Salary    float64
	CompanyID uint
	Company   Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Roles     []Role  `gorm:"many2many:person_roles;"`
}
