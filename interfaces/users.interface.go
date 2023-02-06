package interfaces

import (
	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/libs"
)

type UserRepo interface {
	GetAllUsers() (*models.Users, error)
	GetPageUsers(limit, offset int) (*models.Users, error)
	GetUserById(id string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByToken(token string) (*models.User, error)
	AddUser(data *models.User) (*models.User, error)
	UpdateUser(data *models.User, id string) (*models.User, error)
	DeleteUser(id string) (*models.User, error)
	UpdateToken(id, token string) error
	UserExsist(username string) bool
	EmailExsist(email string) bool
	TokenExsist(token string) bool
}

type UserSrvc interface {
	GetAllUsers() *libs.Response
	GetPageUsers(limit, offset int) *libs.Response
	GetUserById(id string) *libs.Response
	AddUser(data *models.User) *libs.Response
	UpdateUser(data *models.User, id string) *libs.Response
	DeleteUser(id string) *libs.Response
}
