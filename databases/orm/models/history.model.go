package models

import "time"

type History struct {
	HistoryID     uint64      `gorm:"primaryKey" json:"id,omitempty" valid:"-"`
	ReservationID uint64      `gorm:"foreignKey:ReservationID; references:ReservationID" json:"reservation_id" valid:"type(int)"`
	Reservation   Reservation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reservation_data" valid:"-"`
	ReturnStatus  string      `gorm:"default: Not been returned" json:"return_status" valid:"type(string)"`
	CreatedAt     time.Time   `json:"created_at" valid:"-"`
	UpdatedAt     time.Time   `json:"updated_at" valid:"-"`
}

type Histories []History

func (History) TableName() string {
	return "history"
}
