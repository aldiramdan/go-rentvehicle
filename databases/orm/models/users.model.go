package models

import (
	"time"
)

type User struct {
	UserID      uint64    `gorm:"primaryKey" json:"id,omitempty" valid:"-"`
	Username    string    `json:"username,omitempty" valid:"required,type(string)"`
	Email       string    `json:"email,omitempty" valid:"required,email"`
	Role        string    `gorm:"default: user" json:"role,omitempty" valid:"-"`
	Password    string    `json:"password,omitempty" valid:"required,length(8|32)"`
	Name        string    `json:"name,omitempty" valid:"required,type(string)"`
	Gender      string    `json:"gender,omitempty" valid:"type(string)"`
	Address     string    `json:"address,omitempty" valid:"-"`
	Phone       string    `json:"phone,omitempty" valid:"-"`
	BirthDate   string    `json:"birth_date,omitempty" schema:"birth_date" valid:"-"`
	TokenVerify string    `json:"token_verify,omitempty" valid:"-"`
	IsActive    bool      `gorm:"default: false" json:"is_active,omitempty" valid:"-"`
	Picture     string    `json:"picture,omitempty" valid:"-"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	UpdatedAt   time.Time `json:"updated_at" valid:"-"`
}

type Users []User

func (User) TableName() string {
	return "users"
}
