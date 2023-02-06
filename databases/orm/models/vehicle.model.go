package models

import (
	"time"
)

type Vehicle struct {
	VehicleID   string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	Name        string    `json:"name,omitempty" valid:"required,type(string)"`
	Location    string    `json:"location,omitempty" valid:"required,type(string)"`
	Description string    `json:"description,omitempty" valid:"required,type(string)"`
	Price       float64   `json:"price,omitempty" valid:"numeric"`
	Status      string    `json:"status,omitempty" valid:"type(string)"`
	Stock       uint      `json:"stock,omitempty" valid:"numeric"`
	CategoryID  string    `json:"category_id" gorm:"foreignKey:CategoryID; references:CategoryID;" schema:"category_id" valid:"uuidv4" `
	Category    Category  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category_data" valid:"-"`
	Picture     string    `json:"picture,omitempty" valid:"-"`
	Rating      float64   `json:"rating,omitempty" gorm:"default: 0" valid:"type(float64),range(1|5)"`
	TotalRent   int       `json:"total_rent,omitempty" gorm:"default: 0" valid:"-"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	UpdatedAt   time.Time `json:"updated_at" valid:"-"`
}

type Vehicles []Vehicle

func (Vehicle) TableName() string {
	return "vehicle"
}
