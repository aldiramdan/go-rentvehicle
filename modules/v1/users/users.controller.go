package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/interfaces"
	"github.com/aldiramdan/go-backend/libs"
	"github.com/gorilla/mux"
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

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true)
		return
	}

	c.srvc.GetUserById(id).Send(w)

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

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true)
		return
	}

	var data *models.User

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		libs.GetResponse(err.Error(), 500, true)
		return
	}

	c.srvc.UpdateUser(data, id).Send(w)

}

func (c *user_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true)
		return
	}

	c.srvc.DeleteUser(id).Send(w)

}
