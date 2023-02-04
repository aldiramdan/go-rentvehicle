package models

import (
	"time"
)

type Vehicle struct {
	VehicleID   uint64    `gorm:"primaryKey" json:"id,omitempty" valid:"-"`
	Name        string    `json:"name,omitempty" valid:"required,type(string)"`
	Location    string    `json:"location,omitempty" valid:"required,type(string)"`
	Description string    `json:"description,omitempty" valid:"required,type(string)"`
	Price       float64   `json:"price,omitempty" valid:"type(float64)"`
	Status      string    `json:"status,omitempty" valid:"type(string)"`
	Stock       uint      `json:"stock,omitempty" valid:"type(int)"`
	CategoryID  uint64    `gorm:"foreignKey:CategoryID; references:CategoryID;" schema:"category_id" json:"category_id" valid:"type(int)"`
	Category    Category  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category_data" valid:"-"`
	Picture     string    `json:"picture,omitempty" schema:"image" valid:"-"`
	Rating      float64   `gorm:"default: 0" json:"rating" valid:"type(int),length(1|5)"`
	TotalRent   int       `gorm:"default: 0" json:"total_rent" valid:"-"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	UpdatedAt   time.Time `json:"updated_at" valid:"-"`
}

type Vehicles []Vehicle

func (Vehicle) TableName() string {
	return "vehicle"
}
