package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	res "github.com/aldiramdan/go-backend/helpers"
	"github.com/aldiramdan/go-backend/modules/v1/interfaces"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type users_ctrl struct {
	repo interfaces.UserRepo
}

func NewUsersCtrl(repo interfaces.UserRepo) *users_ctrl {

	return &users_ctrl{repo}

}

func (c *users_ctrl) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	result, err := c.repo.GetAllUsers()

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if len(*result) == 0 {
		res.ResponseError(w, http.StatusNotFound, "Data user empty")
		return
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *users_ctrl) GetByUserId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := c.repo.GetByUserId(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.ResponseError(w, http.StatusNotFound, "User not found")
			return
		default:
			res.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *users_ctrl) AddUser(w http.ResponseWriter, r *http.Request) {

	var data models.User

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	salt := 12
	password := []byte(data.Password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, salt)
	data.Password = string(hashedPassword)

	result, err := c.repo.AddUser(&data)

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res.ResponseJson(w, http.StatusCreated, result)

}

func (c *users_ctrl) UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var data models.User

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = c.repo.GetByUserId(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.ResponseError(w, http.StatusNotFound, "User not found")
			return
		default:
			res.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	salt := 12
	password := []byte(data.Password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, salt)
	data.Password = string(hashedPassword)

	result, err := c.repo.UpdateUser(&data, id)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *users_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = c.repo.GetByUserId(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.ResponseError(w, http.StatusNotFound, "User not found")
			return
		default:
			res.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	_, err = c.repo.DeleteUser(id)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := map[string]string{"message": "Users deleted successfully"}

	res.ResponseJson(w, http.StatusOK, response)

}
