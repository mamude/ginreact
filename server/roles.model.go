package main

import "gorm.io/gorm"

// Role model
type Role struct {
	gorm.Model
	Name string
}
