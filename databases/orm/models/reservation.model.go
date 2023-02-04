package models

import "time"

type Reservation struct {
	ReservationID   uint64    `gorm:"primaryKey" json:"id,omitempty" valid:"-"`
	UserID          uint64    `gorm:"foreignKey:UserID; references:UserID" json:"user_id" valid:"type(int)"`
	User            User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_data,omitempty" valid:"-"`
	VehicleID       uint64    `gorm:"foreignKey:VehicleID; references:VehicleID" json:"vehicle_id" valid:"numeric"`
	Vehicle         Vehicle   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"vehicle_data" valid:"-"`
	StartDate       string    `json:"start_date" valid:"-"`
	EndDate         string    `json:"end_date" valid:"-"`
	Quantity        uint      `json:"quantity" valid:"numeric"`
	PaymentCode     string    `json:"payment_code" valid:"-"`
	PaymentMethod   string    `json:"payment_method" valid:"type(string)"`
	PaymentStatus   string    `gorm:"default: Pending" json:"payment_status" valid:"type(string)"`
	Prepayment      float64   `json:"prepayment" valid:"numeric"`
	IsBooked        bool      `gorm:"default: false" json:"is_booked" valid:"-"`
	Rating          float64   `json:"rating" valid:"numeric,length(1|5)"`
	TransactionDate time.Time `gorm:"default: now()" json:"transaction_date" valid:"-"`
	CreatedAt       time.Time `json:"created_at" valid:"-"`
	UpdatedAt       time.Time `json:"updated_at" valid:"-"`
}

type Reservations []Reservation

func (Reservation) TableName() string {
	return "reservation"
}
