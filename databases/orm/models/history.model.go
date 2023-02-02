package models

import "time"

type History struct {
	HistoryID     uint64 `gorm:"primaryKey" json:"id,omitempty"`
	ReservationID uint64 `gorm:"column:ReservationIdForReservation"`
	Reservation   Reservation
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Histories []History

func (History) TableName() string {
	return "history"
}
