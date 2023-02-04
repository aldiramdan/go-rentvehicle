package auth

import (
	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/interfaces"
	"github.com/aldiramdan/go-backend/libs"
)

type auth_service struct {
	repo interfaces.UserRepo
}

type token_response struct {
	Token string `json:"token"`
}

func NewSrvc(repo interfaces.UserRepo) *auth_service {

	return &auth_service{repo}

}

func (s *auth_service) Login(data *models.User) *libs.Response {

	user, err := s.repo.GetByUsername(data.Username)

	if err != nil {
		return libs.GetResponse("user is not exist", 401, true)
	}

	if !libs.CheckPassword(user.Password, data.Password) {
		return libs.GetResponse("password false", 401, true)
	}

	if !user.IsActive {
		return libs.GetResponse("please verify your account", 401, true)
	}

	jwt := libs.NewToken(user.UserID, user.Role)

	token, err := jwt.Create()
	if err != nil {
		return libs.GetResponse(err.Error(), 401, true)
	}

	newToken := "Bearer " + token

	return libs.GetResponse(token_response{Token: newToken}, 200, false)

}

func (s *auth_service) VerifyEmail(token string) *libs.Response {

	if tokenExsist := s.repo.TokenExsist(token); !tokenExsist {
		return libs.GetResponse("failed to verify email", 400, true)
	}

	user, err := s.repo.GetByToken(token)

	if err != nil {
		return libs.GetResponse("user is not exist", 401, true)
	}

	if user.IsActive {
		return libs.GetResponse("already registered, you can login", 401, true)
	}

	var data models.User

	data.IsActive = true

	_, err = s.repo.UpdateUser(&data, user.UserID)

	if err != nil {
		return libs.GetResponse("user is not exist", 401, true)
	}

	response := map[string]string{"message": "successfully verify email "}

	return libs.GetResponse(response, 200, false)

}
