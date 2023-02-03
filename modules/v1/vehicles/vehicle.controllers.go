package vehicles

import (
	"net/http"
	"os"
	"strconv"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/interfaces"
	"github.com/aldiramdan/go-backend/libs"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type vehicle_ctrl struct {
	srvc interfaces.VehicleSrvc
}

func NewCtrl(srvc interfaces.VehicleSrvc) *vehicle_ctrl {

	return &vehicle_ctrl{srvc}

}

func (c *vehicle_ctrl) GetAllVehicles(w http.ResponseWriter, r *http.Request) {

	c.srvc.GetAllVehicles().Send(w)

}

func (c *vehicle_ctrl) GetPopulerVehicle(w http.ResponseWriter, r *http.Request) {

	c.srvc.GetPopulerVehicle().Send(w)

}

func (c *vehicle_ctrl) GetVehicleById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true)
		return
	}

	c.srvc.GetVehicleById(id).Send(w)

}

func (c *vehicle_ctrl) SearchVehicle(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()

	query, ok := vars["s"]
	if !ok {
		libs.GetResponse("Missing query parameter", 400, true)
		return
	}

	c.srvc.SearchVehicle(query[0]).Send(w)

}

func (c *vehicle_ctrl) AddVehicle(w http.ResponseWriter, r *http.Request) {

	var data models.Vehicle

	imageName := r.Context().Value("imageName").(string)
	data.Picture = imageName

	err := schema.NewDecoder().Decode(&data, r.MultipartForm.Value)

	if err != nil {
		_ = os.Remove(imageName)
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.srvc.AddVehicle(&data).Send(w)

}

func (c *vehicle_ctrl) UpdateVehicle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true)
		return
	}

	var data models.Vehicle

	imageName := r.Context().Value("imageName").(string)
	data.Picture = imageName

	err = schema.NewDecoder().Decode(&data, r.MultipartForm.Value)

	if err != nil {
		libs.GetResponse(err.Error(), 500, true)
		return
	}

	c.srvc.UpdateVehicle(&data, id).Send(w)

}

func (c *vehicle_ctrl) DeleteVehicle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true)
		return
	}

	c.srvc.DeleteVehicle(id).Send(w)

}
