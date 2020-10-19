package main

import (
	"gorm.io/gorm"
	"time"
)

// Person model
type Person struct {
	ID        int `gorm:"primaryKey"`
	Age       int
	Level     string
	Country   string
	FirstName string
	LastName  string
	Salary    float64
	CompanyID int
	Company   Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Roles     []Role  `gorm:"many2many:person_roles;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
