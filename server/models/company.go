package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	OpeningYear int
	Name        *string `gorm:"uniqueIndex:idx_name,sort:asc,not null"`
	Description string
}
