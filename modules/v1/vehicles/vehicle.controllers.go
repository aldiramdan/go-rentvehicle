package vehicles

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

type vehicle_ctrl struct {
	repo interfaces.VehicleRepo
}

func NewVehicleCtrl(repo interfaces.VehicleRepo) *vehicle_ctrl {

	return &vehicle_ctrl{repo}

}

func (c *vehicle_ctrl) GetAllVehicles(w http.ResponseWriter, r *http.Request) {

	result, err := c.repo.GetAllVehicles()

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(*result) == 0 {
		res.ResponseError(w, http.StatusNotFound, "Data Vehicle empty")
		return
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *vehicle_ctrl) GetPopulerVehicle(w http.ResponseWriter, r *http.Request) {

	result, err := c.repo.GetPopulerVehicle()

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *vehicle_ctrl) GetVehicleById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := c.repo.GetVehicleById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.ResponseError(w, http.StatusNotFound, "Vehicle not found")
			return
		default:
			res.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *vehicle_ctrl) SearchVehicle(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()

	query, ok := vars["s"]
	if !ok {
		res.ResponseError(w, http.StatusBadRequest, "Missing query parameter")
		return
	}

	result, err := c.repo.SearchVehicle(query[0])

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(*result) == 0 {
		res.ResponseError(w, http.StatusNotFound, "Vehicle not found")
		return
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *vehicle_ctrl) AddVehicle(w http.ResponseWriter, r *http.Request) {

	var data models.Vehicle

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := c.repo.AddVehicle(&data)

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *vehicle_ctrl) UpdateVehicle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var data models.Vehicle

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = c.repo.GetVehicleById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.ResponseError(w, http.StatusNotFound, "Vehicle not found")
			return
		default:
			res.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	result, err := c.repo.UpdateVehicle(&data, id)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *vehicle_ctrl) DeleteVehicle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = c.repo.GetVehicleById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.ResponseError(w, http.StatusNotFound, "Vehicle not found")
			return
		default:
			res.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	_, err = c.repo.DeleteVehicle(id)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := map[string]string{"message": "Vehicle deleted successfully"}

	res.ResponseJson(w, http.StatusOK, response)

}
