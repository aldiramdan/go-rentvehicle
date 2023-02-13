package interfaces

import (
	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/libs"
)

type ReservationRepo interface {
	GetAllReservations() (*models.Reservations, error)
	GetPageReservations(limit, offset int) (*models.Reservations, error)
	GetReservationById(id string) (*models.Reservation, error)
	GetReservationUser(user_id string) (*models.Reservation, error)
	GetReservationByCode(paymentCode string) (*models.Reservation, error)
	AddReservation(data *models.Reservation) (*models.Reservation, error)
	UpdateReservation(data *models.Reservation, paymentCode string) (*models.Reservation, error)
}

type ReservationSrvc interface {
	GetAllReservations() *libs.Response
	GetPageReservations(page, perpage int) *libs.Response
	GetReservationById(id string) *libs.Response
	GetReservationByUser(user_id string) *libs.Response
	GetReservationByCode(paymentCode string) *libs.Response
	AddReservation(data *models.Reservation) *libs.Response
	UpdateReservation(data *models.Reservation, paymentCode string) *libs.Response
}
