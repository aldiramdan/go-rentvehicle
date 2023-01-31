package users

import (
	"errors"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"gorm.io/gorm"
)

type user_repo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *user_repo {

	return &user_repo{db}

}

func (r *user_repo) GetAllUsers() (*models.Users, error) {

	var data models.Users

	if err := r.db.Order("created_at DESC").Find(&data).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	return &data, nil

}

func (r *user_repo) GetByUserId(id uint64) (*models.User, error) {

	var data models.User

	if err := r.db.First(&data, id).Error; err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *user_repo) AddUser(data *models.User) (*models.User, error) {

	if err := r.db.Create(data).Error; err != nil {
		return nil, errors.New("failed to create data")
	}

	return data, nil

}

func (r *user_repo) UpdateUser(data *models.User, id uint64) (*models.User, error) {

	if err := r.db.Model(data).Where("user_id = ?", id).Updates(&data).Error; err != nil {
		return nil, errors.New("failed to update data")
	}

	return data, nil

}

func (r *user_repo) DeleteUser(id uint64) (*models.User, error) {

	var data models.User

	if err := r.db.Delete(data, id).Error; err != nil {
		return nil, errors.New("failed to delete data")
	}

	return &data, nil

}
