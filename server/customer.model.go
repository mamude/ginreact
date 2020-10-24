package main

import (
	"gorm.io/gorm"
	"time"
)

// Customer model
type Customer struct {
	ID        int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Phone     string
	Email     *string `gorm:"uniqueIndex:idx_customer_name,sort:asc,not null"`
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
