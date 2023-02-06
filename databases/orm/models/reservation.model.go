package models

import "time"

type Reservation struct {
	ReservationID   string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	UserID          string    `gorm:"foreignKey:UserID; references:UserID" json:"user_id" valid:"uuidv4"`
	User            User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_data,omitempty" valid:"-"`
	VehicleID       string    `gorm:"foreignKey:VehicleID; references:VehicleID" json:"vehicle_id" valid:"uuidv4"`
	Vehicle         Vehicle   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"vehicle_data" valid:"-"`
	StartDate       string    `json:"start_date" valid:"-"`
	EndDate         string    `json:"end_date" valid:"-"`
	Quantity        uint      `json:"quantity" valid:"numeric"`
	PaymentCode     string    `json:"payment_code" valid:"-"`
	PaymentMethod   string    `json:"payment_method" valid:"type(string)"`
	PaymentStatus   string    `gorm:"default: Pending" json:"payment_status" valid:"type(string)"`
	Prepayment      float64   `json:"prepayment" valid:"numeric"`
	IsBooked        bool      `gorm:"default: false" json:"is_booked" valid:"-"`
	Rating          float64   `json:"rating" valid:"type(float64),range(1|5)"`
	TransactionDate time.Time `gorm:"default: now()" json:"transaction_date" valid:"-"`
	CreatedAt       time.Time `json:"created_at" valid:"-"`
	UpdatedAt       time.Time `json:"updated_at" valid:"-"`
}

type Reservations []Reservation

func (Reservation) TableName() string {
	return "reservation"
}
