package interfaces

import (
	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/libs"
)

type CategoryRepo interface {
	GetAllCategories() (*models.Categories, error)
	GetPageCategories(limit, offset int) (*models.Categories, error)
	GetCategoryById(id string) (*models.Category, error)
	SearchCategories(query string) (*models.Categories, error)
	AddCategory(data *models.Category) (*models.Category, error)
	UpdateCategory(data *models.Category, id string) (*models.Category, error)
	DeleteCategory(id string) (*models.Category, error)
}

type CategorySrvc interface {
	GetAllCategories() *libs.Response
	GetPageCategories(page, perpage int) *libs.Response
	GetCategoryById(id string) *libs.Response
	SearchCategories(query string) *libs.Response
	AddCategory(data *models.Category) *libs.Response
	UpdateCategory(data *models.Category, id string) *libs.Response
	DeleteCategory(id string) *libs.Response
}
