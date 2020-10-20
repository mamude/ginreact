package main

import (
	"gorm.io/gorm"
	"time"
)

// Role model
type Role struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}