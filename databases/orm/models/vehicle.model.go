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
	CategoryID  uint64    `gorm:"foreignKey:CategoryID; references:CategoryID" schema:"category_id" json:"category_id"`
	Category    Category  `json:"category_data"`
	Picture     string    `json:"picture" schema:"image"`
	Rating      float64   `gorm:"default: 0" json:"rating"`
	TotalRent   int       `gorm:"default: 0" json:"total_rent"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Vehicles []Vehicle

func (Vehicle) TableName() string {
	return "vehicle"
}
