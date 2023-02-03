package models

import "time"

type Category struct {
	CategoryID uint64    `gorm:"primaryKey" json:"id,omitempty" valid:"-"`
	Name       string    `json:"name" valid:"required,type(string)"`
	CreatedAt  time.Time `json:"created_at" valid:"-"`
	UpdatedAt  time.Time `json:"updated_at" valid:"-"`
}

type Categories []Category

func (Category) TableName() string {
	return "category"
}
