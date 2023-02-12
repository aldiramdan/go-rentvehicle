package reservations

import (
	"fmt"

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
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *reservation_srvc) GetReservationById(id string) *libs.Response {

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

func (s *reservation_srvc) GetReservationByCode(paymentCode string) *libs.Response {

	result, err := s.repo.GetReservationByCode(paymentCode)

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
		return libs.GetResponse(err.Error(), 500, true)
	}

	if data.PaymentMethod == "Cash" {
		data.IsBooked = true
		data.PaymentStatus = "Paid"
	}

	data.PaymentCode = fmt.Sprintf("%s%s", "GRV-", paymentCode)

	result, err := s.repo.AddReservation(data)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *reservation_srvc) UpdateReservation(data *models.Reservation, paymentCode string) *libs.Response {

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
		return libs.GetResponse("payment successfully", 401, true)
	} else {
		data.PaymentStatus = "Paid"
		data.IsBooked = true
	}

	result, err := s.repo.UpdateReservation(data, paymentCode)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}
