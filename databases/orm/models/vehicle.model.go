package models

import (
	"time"
)

type Vehicle struct {
	VehicleID   uint64    `gorm:"primaryKey" json:"id,omitempty"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Status      string    `json:"status"`
	Stock       uint      `json:"stock"`
	CategoryID  uint64    `gorm:"unique" json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryID; references:CategoryID;" json:"category"`
	Picture     string    `json:"picture"`
	Rating      float64   `gorm:"default: 0" json:"rating"`
	TotalRent   int       `gorm:"default: 0" json:"total_rent"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Vehicles []Vehicle

func (Vehicle) TableName() string {
	return "vehicle"
}
