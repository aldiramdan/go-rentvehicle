package models

import (
	"time"
)

type User struct {
	UserID    uint64    `gorm:"primaryKey" json:"id,omitempty"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	BirthDate string    `json:"birth_date"`
	Picture   string    `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

func (User) TableName() string {
	return "users"
}
