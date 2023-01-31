package interfaces

import "github.com/aldiramdan/go-backend/databases/orm/models"

type UserRepo interface {
	GetAllUsers() (*models.Users, error)
	GetByUserId(id uint64) (*models.User, error)
	AddUser(data *models.User) (*models.User, error)
	UpdateUser(data *models.User, id uint64) (*models.User, error)
	DeleteUser(id uint64) (*models.User, error)
}
