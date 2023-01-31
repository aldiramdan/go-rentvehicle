package reservations

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	res "github.com/aldiramdan/go-backend/helpers"
	"github.com/aldiramdan/go-backend/modules/v1/interfaces"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type reservation_ctrl struct {
	repo interfaces.ReservationRepo
}

func NewReservationCtrl(repo interfaces.ReservationRepo) *reservation_ctrl {

	return &reservation_ctrl{repo}

}

func (c *reservation_ctrl) GetAllReservations(w http.ResponseWriter, r *http.Request) {

	result, err := c.repo.GetAllReservations()

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if len(*result) == 0 {
		res.ResponseError(w, http.StatusNotFound, "Data transaction empty")
		return
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *reservation_ctrl) GetReservationById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := c.repo.GetReservationById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.ResponseError(w, http.StatusNotFound, "Transaction not found")
			return
		default:
			res.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *reservation_ctrl) AddReservation(w http.ResponseWriter, r *http.Request) {

	var data models.Reservation

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := c.repo.AddReservation(&data)

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res.ResponseJson(w, http.StatusCreated, result)

}

func (c *reservation_ctrl) Payment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var data models.Reservation

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := c.repo.Payment(&data, id)

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res.ResponseJson(w, http.StatusCreated, result)

}
