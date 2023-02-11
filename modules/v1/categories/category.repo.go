package categories

import (
	"errors"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"gorm.io/gorm"
)

type category_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *category_repo {

	return &category_repo{db}
	
}

func (r *category_repo) GetAllCategories() (*models.Categories, error) {

	var data models.Categories

	if err := r.db.
		Order("created_at DESC").
		Find(&data).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	if len(data) == 0 {
		return nil, errors.New("data category is empty")
	}

	return &data, nil

}

func (r *category_repo) GetPageCategories(limit, offset int) (*models.Categories, error) {

	var data models.Categories

	if err := r.db.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&data).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	if len(data) == 0 {
		return nil, errors.New("data category is empty")
	}

	return &data, nil

}

func (r *category_repo) GetCategoryById(id string) (*models.Category, error) {

	var data models.Category

	if err := r.db.
		First(&data, "category_id = ?", id).Error; err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *category_repo) SearchCategories(query string) (*models.Categories, error) {

	var data models.Categories

	if err := r.db.
		Order("created_at DESC").
		Where("name LIKE ? ", "%"+query+"%").
		Find(&data).Error; err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("search data category is empty")
	}

	return &data, nil

}

func (r *category_repo) AddCategory(data *models.Category) (*models.Category, error) {

	if err := r.db.
		Create(data).Error; err != nil {
		return nil, errors.New("failed to create data")
	}

	return data, nil

}

func (r *category_repo) UpdateCategory(data *models.Category, id string) (*models.Category, error) {

	if err := r.db.
		Model(&data).
		Where("category_id = ?", id).
		Updates(&data).Error; err != nil {
		return nil, errors.New("failed to update data")
	}

	return data, nil

}

func (r *category_repo) DeleteCategory(id string) (*models.Category, error) {

	var data models.Category

	if err := r.db.
		Delete(data, "category_id = ?", id).Error; err != nil {
		return nil, err
	}

	return &data, nil

}
