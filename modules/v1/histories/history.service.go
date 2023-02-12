package histories

import (
	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/interfaces"
	"github.com/aldiramdan/go-backend/libs"
	"gorm.io/gorm"
)

type history_service struct {
	repo interfaces.HistoryRepo
}

func NewSrvc(repo interfaces.HistoryRepo) *history_service {

	return &history_service{repo}

}

func (s *history_service) GetAllHistories(user_id string) *libs.Response {

	result, err := s.repo.GetAllHistories(user_id)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *history_service) GetPageHistories(user_id string, page, perpage int) *libs.Response {

	offset := (page - 1) * perpage

	result, err := s.repo.GetPageHistories(user_id, perpage, offset)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *history_service) GetHistoryById(id string) *libs.Response {

	result, err := s.repo.GetHistoryById(id)

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

func (s *history_service) SearchHistory(user_id string, query string) *libs.Response {

	result, err := s.repo.SearchHistory(user_id, query)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *history_service) AddHistory(data *models.History) *libs.Response {

	result, err := s.repo.AddHistory(data)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *history_service) UpdateHistory(data *models.History, id string) *libs.Response {

	datas, err := s.repo.GetHistoryById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	if datas.Reservation.PaymentStatus == "Pending" {
		return libs.GetResponse("you need to pay", 401, true)
	}

	result, err := s.repo.UpdateHistory(data, id)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *history_service) DeleteHistory(id string) *libs.Response {

	_, err := s.repo.GetHistoryById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	_, err = s.repo.DeleteHistory(id)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	response := map[string]string{"message": "History deleted successfully"}

	return libs.GetResponse(response, 200, false)

}
