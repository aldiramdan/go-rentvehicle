package users

import (
	"errors"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"gorm.io/gorm"
)

type user_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *user_repo {

	return &user_repo{db}

}

func (r *user_repo) GetAllUsers() (*models.Users, error) {

	var data models.Users

	if err := r.db.
		Select("user_id, username, email, name, gender, address, phone, birth_date, picture").
		Order("created_at DESC").
		Find(&data).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	if len(data) == 0 {
		return nil, errors.New("data user is empty")
	}

	return &data, nil

}

func (r *user_repo) GetPageUsers(limit, offset int) (*models.Users, error) {

	var data models.Users

	if err := r.db.
		Select("user_id, username, email, name, gender, address, phone, birth_date, picture, created_at, updated_at").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&data).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	if len(data) == 0 {
		return nil, errors.New("data user is empty")
	}

	return &data, nil

}

func (r *user_repo) GetUserById(id string) (*models.User, error) {

	var data models.User

	if err := r.db.
		Select("user_id, username, email, name, gender, address, phone, birth_date, picture, created_at, updated_at").
		Find(&data, "user_id = ?", id).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	return &data, nil

}

func (r *user_repo) GetByUsername(username string) (*models.User, error) {

	var data models.User

	if err := r.db.
		First(&data, "username = ?", username).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	return &data, nil

}

func (r *user_repo) GetByEmail(email string) (*models.User, error) {

	var data models.User

	if err := r.db.
		First(&data, "email = ?", email).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	return &data, nil

}

func (r *user_repo) GetByToken(token string) (*models.User, error) {

	var data models.User

	if err := r.db.
		First(&data, "token_verify = ?", token).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	return &data, nil

}

func (r *user_repo) AddUser(data *models.User) (*models.User, error) {

	if err := r.db.
		Create(data).Error; err != nil {
		return nil, errors.New("failed to create data")
	}

	return data, nil

}

func (r *user_repo) UpdateUser(data *models.User, id string) (*models.User, error) {

	if err := r.db.
		Model(data).
		Where("user_id = ?", id).
		Updates(&data).Error; err != nil {
		return nil, errors.New("failed to update data")
	}

	return data, nil

}

func (r *user_repo) DeleteUser(id string) (*models.User, error) {

	var data models.User

	if err := r.db.
		Delete(data, "user_id = ?", id).Error; err != nil {
		return nil, errors.New("failed to delete data")
	}

	return &data, nil

}

func (r *user_repo) UpdateToken(id, token string) error {

	var data models.User

	if err := r.db.
		Model(data).
		Where("user_id = ?", id).
		Update("token_verify", token).Error; err != nil {
		return errors.New("failed to update data")
	}

	return nil

}

func (r *user_repo) UserExists(username string) bool {

	var data models.User

	err := r.db.
		First(&data, "username = ?", username)

	return err.Error == nil

}

func (r *user_repo) EmailExists(email string) bool {

	var data models.User

	err := r.db.
		First(&data, "email = ?", email)

	return err.Error == nil

}

func (r *user_repo) TokenExists(token string) bool {

	var data models.User

	err := r.db.
		First(&data, "token_verify = ?", token)

	return err.Error == nil

}
