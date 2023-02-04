package users

import (
	"net/http"
	"os"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/interfaces"
	"github.com/aldiramdan/go-backend/libs"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema"
)

type user_ctrl struct {
	srvc interfaces.UserSrvc
}

func NewCtrl(srvc interfaces.UserSrvc) *user_ctrl {

	return &user_ctrl{srvc}

}

func (c *user_ctrl) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	c.srvc.GetAllUsers().Send(w)

}

func (c *user_ctrl) GetUserById(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user")

	c.srvc.GetUserById(user_id.(uint64)).Send(w)

}

func (c *user_ctrl) AddUser(w http.ResponseWriter, r *http.Request) {

	var data models.User

	imageName := r.Context().Value("imageName").(string)
	data.Picture = imageName

	err := schema.NewDecoder().Decode(&data, r.MultipartForm.Value)

	if err != nil {
		_ = os.Remove(imageName)
		libs.GetResponse(err.Error(), 500, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(data)

	if err != nil {
		_ = os.Remove(imageName)
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.srvc.AddUser(&data).Send(w)

}

func (c *user_ctrl) UpdateUser(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user")

	var data models.User

	imageName := r.Context().Value("imageName").(string)
	data.Picture = imageName

	err := schema.NewDecoder().Decode(&data, r.MultipartForm.Value)

	if err != nil {
		_ = os.Remove(imageName)
		libs.GetResponse(err.Error(), 500, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(data)

	if err != nil {
		_ = os.Remove(imageName)
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.srvc.UpdateUser(&data, user_id.(uint64)).Send(w)

}

func (c *user_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user")

	c.srvc.DeleteUser(user_id.(uint64)).Send(w)

}
