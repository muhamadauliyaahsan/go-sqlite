package model

import "time"

// Product represents a product entity
// @Description Product information
type Product struct {
	ID          uint      `json:"id,omitempty" gorm:"primaryKey"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Price       float64   `json:"price" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}
