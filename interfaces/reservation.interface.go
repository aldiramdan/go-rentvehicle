package interfaces

import (
	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/libs"
)

type ReservationRepo interface {
	GetAllReservations() (*models.Reservations, error)
	GetReservationById(id uint64) (*models.Reservation, error)
	AddReservation(data *models.Reservation) (*models.Reservation, error)
	Payment(data *models.Reservation, paymentCode string) (*models.Reservation, error)
}

type ReservationSrvc interface {
	GetAllReservations() *libs.Response
	GetReservationById(id uint64) *libs.Response
	AddReservation(data *models.Reservation) *libs.Response
	Payment(data *models.Reservation, paymentCode string) *libs.Response
}
