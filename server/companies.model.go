package main

import "gorm.io/gorm"

// Company model
type Company struct {
	gorm.Model
	OpeningYear int
	Name        *string `gorm:"uniqueIndex:idx_name,sort:asc,not null"`
	Description string
}
