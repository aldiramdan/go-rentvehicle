package interfaces

import "github.com/aldiramdan/go-backend/databases/orm/models"

type CategoryRepo interface {
	GetAllCategories() (*models.Categories, error)
	GetCategoryById(id uint64) (*models.Category, error)
	SearchCategories(query string) (*models.Categories, error)
	AddCategory(data *models.Category) (*models.Category, error)
	UpdateCategory(data *models.Category, id uint64) (*models.Category, error)
	DeleteCategory(id uint64) (*models.Category, error)
}
