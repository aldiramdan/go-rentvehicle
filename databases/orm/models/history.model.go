package models

import "time"

type History struct {
	HistoryID     string      `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	ReservationID string      `gorm:"foreignKey:ReservationID; references:ReservationID" json:"reservation_id" valid:"uuidv4"`
	Reservation   Reservation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reservation_data" valid:"-"`
	ReturnStatus  string      `json:"return_status" gorm:"default: Not been returned" valid:"type(string)"`
	CreatedAt     time.Time   `json:"created_at" valid:"-"`
	UpdatedAt     time.Time   `json:"updated_at" valid:"-"`
}

type Histories []History

func (History) TableName() string {
	return "history"
}
