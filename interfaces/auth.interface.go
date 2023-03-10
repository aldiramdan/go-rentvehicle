package interfaces

import (
	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/libs"
)

type AuthSrvc interface {
	Login(data *models.User) *libs.Response
	VerifyEmail(token string) *libs.Response
	ResendEmail(data *models.User) *libs.Response
}
