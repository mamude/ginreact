package main

import (
	"gorm.io/gorm"
	"time"
)

// Order model
type Order struct {
	ID         int `gorm:"primaryKey"`
	MarketID   int
	Market     Market `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CustomerID int
	Customer   Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status     string
	Total      float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// OrderItem model
type OrderItem struct {
	ID        int `gorm:"primaryKey"`
	OrderID   int
	Order     Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductID int
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amount    int
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
