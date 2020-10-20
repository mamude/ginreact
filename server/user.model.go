package main

import (
	"gorm.io/gorm"
	"time"
)

// User model
type User struct {
	ID        int `gorm:"primaryKey"`
	Username  string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
