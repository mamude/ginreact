package models

import (
	"gorm.io/gorm"
	"time"
)

// CategoryBusiness model
type CategoryBusiness struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Category model
type Category struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

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

// Market model
type Market struct {
	ID                 int `gorm:"primaryKey"`
	CategoryBusinessID int
	CategoryBusiness   CategoryBusiness `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name               *string          `gorm:"uniqueIndex:idx_market_name,sort:asc,not null"`
	Description        string
	URL                string
	Email              string
	Phone              string
	Rating             float64
	Distance           float64
	DeliveryTime       string
	DeliveryTax        float64
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	Products           []Product
}

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

// Product model
type Product struct {
	ID          int `gorm:"primaryKey"`
	MarketID    int
	Market      Market `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CategoryID  int
	Category    Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string
	Description string
	Price       float64
	Amount      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// ShoppingCart model
type ShoppingCart struct {
	ID         int `gorm:"primaryKey"`
	CustomerID int
	Customer   Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MarketID   int
	Market     Market `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductID  int
	Product    Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Price      float64
	Amount     int
	SessionID  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

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
