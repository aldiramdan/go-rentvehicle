package models

import "time"

type Reservation struct {
	ReservationID   uint64 `gorm:"primaryKey" json:"id,omitempty"`
	UserID          uint64 `gorm:"column:UserIdForUser"`
	User            User
	VehicleID       uint64 `gorm:"column:VehiclesIdForVehicles"`
	Vehicle         Vehicle
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

type Reservations []Reservation

func (Reservation) TableName() string {
	return "reservation"
}
