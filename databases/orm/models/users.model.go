package models

import (
	"time"
)

type User struct {
	UserID    uint64    `gorm:"primaryKey" json:"id,omitempty" valid:"-"`
	Username  string    `json:"username" valid:"required,type(string)"`
	Email     string    `json:"email" valid:"required,email"`
	Role      string    `gorm:"default: user" json:"role,omitempty" valid:"-"`
	Password  string    `json:"password,omitempty" valid:"required,length(8|32)"`
	Name      string    `json:"name" valid:"required,alpha"`
	Gender    string    `json:"gender" valid:"type(string)"`
	Address   string    `json:"address" valid:"-"`
	Phone     string    `json:"phone" valid:"-"`
	BirthDate string    `json:"birth_date" schema:"birth_date" valid:"-"`
	Picture   string    `json:"picture" schema:"image" valid:"-"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

type Users []User

func (User) TableName() string {
	return "users"
}
