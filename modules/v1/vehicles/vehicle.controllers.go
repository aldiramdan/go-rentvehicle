package vehicles

import (
	"net/http"
	"os"
	"strconv"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/interfaces"
	"github.com/aldiramdan/go-backend/libs"
	"github.com/asaskevich/govalidator"
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

func (c *vehicle_ctrl) GetPageVehicles(w http.ResponseWriter, r *http.Request) {

	varsPage := r.URL.Query().Get("page")
	varsPerPage := r.URL.Query().Get("perpage")

	var page, perpage int
	var err error

	if varsPage != "" {
		page, err = strconv.Atoi(varsPage)
		if err != nil {
			libs.GetResponse(err.Error(), 400, true).Send(w)
			return
		}
	} else {
		page = 1
	}

	if varsPerPage != "" {
		perpage, err = strconv.Atoi(varsPerPage)
		if err != nil {
			libs.GetResponse(err.Error(), 400, true).Send(w)
			return
		}
	} else {
		perpage = 5
	}

	c.srvc.GetPageVehicles(page, perpage).Send(w)

}

func (c *vehicle_ctrl) GetVehicleById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	c.srvc.GetVehicleById(id).Send(w)

}

func (c *vehicle_ctrl) SearchVehicle(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()

	query, ok := vars["s"]
	if !ok {
		libs.GetResponse("Missing query parameter", 400, true).Send(w)
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

	_, err = govalidator.ValidateStruct(data)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.srvc.AddVehicle(&data).Send(w)

}

func (c *vehicle_ctrl) UpdateVehicle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	var data models.Vehicle

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
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.srvc.UpdateVehicle(&data, id).Send(w)

}

func (c *vehicle_ctrl) DeleteVehicle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	c.srvc.DeleteVehicle(id).Send(w)

}
