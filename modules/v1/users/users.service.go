package users

import (
	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/interfaces"
	"github.com/aldiramdan/go-backend/libs"
	"gorm.io/gorm"
)

type user_service struct {
	repo interfaces.UserRepo
}

func NewSrvc(repo interfaces.UserRepo) *user_service {

	return &user_service{repo}

}

func (s *user_service) GetAllUsers() *libs.Response {

	data, err := s.repo.GetAllUsers()

	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	return libs.GetResponse(data, 200, false)

}

func (s *user_service) GetById(id uint64) *libs.Response {

	data, err := s.repo.GetById(id)

	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	return libs.GetResponse(data, 200, false)

}

func (s *user_service) AddUser(data *models.User) *libs.Response {

	if userExsist := s.repo.UserExsist(data.Username); userExsist {
		return libs.GetResponse("username is already registered", 400, true)
	}

	if emailExsist := s.repo.EmailExsist(data.Email); emailExsist {
		return libs.GetResponse("email is already registered", 400, true)
	}

	hashPassword, err := libs.HashPassword(data.Password)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	data.Password = hashPassword
	result, err := s.repo.AddUser(data)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *user_service) UpdateUser(data *models.User, id uint64) *libs.Response {

	_, err := s.repo.GetById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	hashPassword, err := libs.HashPassword(data.Password)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	data.Password = hashPassword
	result, err := s.repo.UpdateUser(data, id)

	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *user_service) DeleteUser(id uint64) *libs.Response {

	_, err := s.repo.GetById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	_, err = s.repo.DeleteUser(id)

	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	response := map[string]string{"message": "User deleted successfully"}

	return libs.GetResponse(response, 200, false)

}
