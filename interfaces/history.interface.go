package interfaces

import (
	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/libs"
)

type HistoryRepo interface {
	GetAllHistories(user_id uint64) (*models.Histories, error)
	GetHistoryById(id uint64) (*models.History, error)
	SearchHistory(user_id uint64, query string) (*models.Histories, error)
	AddHistory(data *models.History) (*models.History, error)
	UpdateHistory(data *models.History, id uint64) (*models.History, error)
	DeleteHistory(id uint64) (*models.History, error)
}

type HistorySrvc interface {
	GetAllHistories(user_id uint64) *libs.Response
	GetHistoryById(id uint64) *libs.Response
	SearchHistory(user_id uint64, query string) *libs.Response
	AddHistory(data *models.History) *libs.Response
	UpdateHistory(data *models.History, id uint64) *libs.Response
	DeleteHistory(id uint64) *libs.Response
}
