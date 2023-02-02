package interfaces

import "github.com/aldiramdan/go-backend/databases/orm/models"

type HistoryRepo interface {
	GetAllHistories() (*models.Histories, error)
	GetHistoryById(id uint64) (*models.History, error)
	SearchHistory(query string) (*models.Histories, error)
	AddHistory(data *models.History) (*models.History, error)
	UpdateHistory(data *models.History, id uint64) (*models.History, error)
	DeleteHistory(id uint64) (*models.History, error)
}
