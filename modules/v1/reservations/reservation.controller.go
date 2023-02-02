package reservations

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/interfaces"
	"github.com/aldiramdan/go-backend/libs"
	"github.com/gorilla/mux"
)

type reservation_ctrl struct {
	srvc interfaces.ReservationSrvc
}

func NewCtrl(srvc interfaces.ReservationSrvc) *reservation_ctrl {

	return &reservation_ctrl{srvc}

}

func (c *reservation_ctrl) GetAllReservations(w http.ResponseWriter, r *http.Request) {

	c.srvc.GetAllReservations().Send(w)

}

func (c *reservation_ctrl) GetReservationById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true)
		return
	}

	c.srvc.GetReservationById(id).Send(w)

}

func (c *reservation_ctrl) AddReservation(w http.ResponseWriter, r *http.Request) {

	var data *models.Reservation

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true)
		return
	}

	c.srvc.AddReservation(data).Send(w)

}

func (c *reservation_ctrl) Payment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true)
		return
	}

	var data *models.Reservation

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		libs.GetResponse(err.Error(), 500, true)
		return
	}

	c.srvc.Payment(data, id).Send(w)

}
