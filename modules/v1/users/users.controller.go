package users

import (
	"net/http"
	"os"
	"strconv"

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

func (c *user_ctrl) GetPageUsers(w http.ResponseWriter, r *http.Request) {

	varsPage := r.URL.Query()

	qPage, ok := varsPage["page"]
	if !ok {
		libs.GetResponse("page is required", 400, true).Send(w)
		return
	}

	varsPerPage := r.URL.Query()

	qPerpage, ok := varsPerPage["perpage"]
	if !ok {
		libs.GetResponse("perpage is required", 400, true).Send(w)
		return
	}

	var page, perpage int
	var err error

	if qPage[0] != "" {
		page, err = strconv.Atoi(qPage[0])
		if err != nil {
			libs.GetResponse(err.Error(), 400, true).Send(w)
			return
		}
	} else {
		page = 1
	}

	if qPerpage[0] != "" {
		perpage, err = strconv.Atoi(qPerpage[0])
		if err != nil {
			libs.GetResponse(err.Error(), 400, true).Send(w)
			return
		}
	} else {
		perpage = 5
	}

	c.srvc.GetPageUsers(page, perpage).Send(w)

}

func (c *user_ctrl) GetUserById(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user")

	c.srvc.GetUserById(user_id.(string)).Send(w)

}

func (c *user_ctrl) AddUser(w http.ResponseWriter, r *http.Request) {

	var data models.User

	imageName := r.Context().Value("imageName").(string)
	data.Picture = imageName

	err := schema.NewDecoder().Decode(&data, r.MultipartForm.Value)

	if err != nil {
		_ = os.Remove(imageName)
		libs.GetResponse(err.Error(), 400, true).Send(w)
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
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(data)

	if err != nil {
		_ = os.Remove(imageName)
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.srvc.UpdateUser(&data, user_id.(string)).Send(w)

}

func (c *user_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user")

	c.srvc.DeleteUser(user_id.(string)).Send(w)

}
