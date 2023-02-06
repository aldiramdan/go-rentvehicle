package reservations

import (
	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/interfaces"
	"github.com/aldiramdan/go-backend/libs"
	"gorm.io/gorm"
)

type reservation_srvc struct {
	repo interfaces.ReservationRepo
}

func NewSrvc(repo interfaces.ReservationRepo) *reservation_srvc {

	return &reservation_srvc{repo}

}

func (s *reservation_srvc) GetAllReservations() *libs.Response {

	result, err := s.repo.GetAllReservations()

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *reservation_srvc) GetPageReservations(page, perpage int) *libs.Response {

	offset := (page - 1) * perpage

	result, err := s.repo.GetPageReservations(perpage, offset)

	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *reservation_srvc) GetReservationById(id uint64) *libs.Response {

	result, err := s.repo.GetReservationById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	return libs.GetResponse(result, 200, false)

}

func (s *reservation_srvc) AddReservation(data *models.Reservation) *libs.Response {

	paymentCode, err := libs.CodeCrypt(6)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	if data.PaymentMethod == "Cash" {
		data.IsBooked = true
		data.PaymentStatus = "Paid"
	}

	data.PaymentCode = paymentCode

	result, err := s.repo.AddReservation(data)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *reservation_srvc) Payment(data *models.Reservation, paymentCode string) *libs.Response {

	datas, err := s.repo.GetReservationByCode(paymentCode)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	if datas.PaymentStatus == "Paid" {
		return libs.GetResponse("Payment successfully", 400, true)
	}

	data.PaymentStatus = "Paid"
	data.IsBooked = true

	result, err := s.repo.Payment(data, paymentCode)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}
