package models

import "time"

type Category struct {
	CategoryID uint64    `gorm:"primaryKey" json:"id,omitempty"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Categories []Category

func (Category) TableName() string {
	return "category"
}
