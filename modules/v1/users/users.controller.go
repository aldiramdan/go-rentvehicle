package users

import (
	"encoding/json"
	"net/http"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/interfaces"
	"github.com/aldiramdan/go-backend/libs"
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

	var data *models.User

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		libs.GetResponse(err.Error(), 500, true)
		return
	}

	c.srvc.AddUser(data).Send(w)

}

func (c *user_ctrl) UpdateUser(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user")

	var data *models.User

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		libs.GetResponse(err.Error(), 500, true)
		return
	}

	c.srvc.UpdateUser(data, user_id.(uint64)).Send(w)

}

func (c *user_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user")

	c.srvc.DeleteUser(user_id.(uint64)).Send(w)

}
