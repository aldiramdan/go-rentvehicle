package models

import "time"

type History struct {
	HistoryID       uint64    `gorm:"primaryKey" json:"history_id,omitempty"`
	UserID          uint      `json:"users_id"`
	User            User      `gorm:"foreignKey:UserID; references:UserID"`
	VehicleID       uint      `json:"vehicle_id"`
	Vehicle         Vehicle   `gorm:"foreignKey:VehicleID; references:VehicleID"`
	StartDate       string    `json:"start_date"`
	EndDate         string    `json:"end_date"`
	Quantity        uint      `json:"quantity"`
	PaymentCode     string    `json:"payment_code"`
	PaymentMethod   string    `json:"payment_method"`
	PaymentStatus   string    `gorm:"default: Pending" json:"payment_status"`
	Prepayment      float64   `json:"prepayment"`
	IsBooked        bool      `gorm:"default: false" json:"is_booked"`
	ReturnStatus    string    `gorm:"default: Not been returned" json:"return_status"`
	Rating          float64   `json:"rating"`
	TransactionDate time.Time `gorm:"default: now()" json:"transaction_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Histories []History

func (History) TableName() string {
	return "history"
}
