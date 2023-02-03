package users

import (
	"os"

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

	result, err := s.repo.GetAllUsers()

	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *user_service) GetUserById(id uint64) *libs.Response {

	result, err := s.repo.GetUserById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	return libs.GetResponse(result, 200, false)

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

	tokenVeify, err := libs.CodeCrypto()
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	data.TokenVerify = tokenVeify

	link := os.Getenv("BASE_URL") + "/auth/confirm_email/" + tokenVeify
	err = libs.SendMail(data.Email, "Veify Email", link)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	result, err := s.repo.AddUser(data)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *user_service) UpdateUser(data *models.User, id uint64) *libs.Response {

	datas, err := s.repo.GetUserById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	if data.Password != "" {
		hashPassword, err := libs.HashPassword(data.Password)
		if err != nil {
			return libs.GetResponse(err.Error(), 400, true)
		}
		data.Password = hashPassword
	}

	if datas.Picture != "public/default_image.jpg" {
		_ = os.Remove(datas.Picture)
	}

	result, err := s.repo.UpdateUser(data, id)

	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *user_service) DeleteUser(id uint64) *libs.Response {

	data, err := s.repo.GetUserById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	if data.Picture != "public/default_image.jpg" {
		_ = os.Remove(data.Picture)
	}

	_, err = s.repo.DeleteUser(id)

	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	response := map[string]string{"message": "User deleted successfully"}

	return libs.GetResponse(response, 200, false)

}
