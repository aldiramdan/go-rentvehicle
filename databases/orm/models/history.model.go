package models

import "time"

type History struct {
	HistoryID     uint64      `gorm:"primaryKey" json:"id,omitempty"`
	ReservationID uint64      `gorm:"unique" json:"reservation_id"`
	Reservation   Reservation `gorm:"foreignKey:ReservationID;references:ReservationID;" json:"reservation,omitempty"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

type Histories []History

func (History) TableName() string {
	return "history"
}
