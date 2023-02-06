package interfaces

import (
	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/libs"
)

type HistoryRepo interface {
	GetAllHistories(user_id string) (*models.Histories, error)
	GetPageHistories(user_id string, limit, offset int) (*models.Histories, error)
	GetHistoryById(id string) (*models.History, error)
	SearchHistory(user_id string, query string) (*models.Histories, error)
	AddHistory(data *models.History) (*models.History, error)
	UpdateHistory(data *models.History, id string) (*models.History, error)
	DeleteHistory(id string) (*models.History, error)
}

type HistorySrvc interface {
	GetAllHistories(user_id string) *libs.Response
	GetPageHistories(user_id string, page, perpage int) *libs.Response
	GetHistoryById(id string) *libs.Response
	SearchHistory(user_id string, query string) *libs.Response
	AddHistory(data *models.History) *libs.Response
	UpdateHistory(data *models.History, id string) *libs.Response
	DeleteHistory(id string) *libs.Response
}
