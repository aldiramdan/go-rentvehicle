package interfaces

import "github.com/aldiramdan/go-backend/databases/orm/models"

type ReservationRepo interface {
	GetAllReservations() (*models.Reservations, error)
	GetReservationById(id uint64) (*models.Reservation, error)
	AddReservation(data *models.Reservation) (*models.Reservation, error)
	Payment(data *models.Reservation, id uint64) (*models.Reservation, error)
}
