package reservations

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/interfaces"
	"github.com/aldiramdan/go-backend/libs"
	"github.com/asaskevich/govalidator"
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

func (c *reservation_ctrl) GetPageReservations(w http.ResponseWriter, r *http.Request) {

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

	c.srvc.GetPageReservations(page, perpage).Send(w)

}

func (c *reservation_ctrl) GetReservationById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	c.srvc.GetReservationById(id).Send(w)

}

func (c *reservation_ctrl) GetReservationByCode(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	paymentCode, ok := vars["payment_code"]

	if !ok {
		libs.GetResponse("Payment Code not found in request", 400, true).Send(w)
		return
	}

	c.srvc.GetReservationByCode(paymentCode).Send(w)

}

func (c *reservation_ctrl) AddReservation(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user")

	var data *models.Reservation

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(data)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	data.UserID = user_id.(string)

	c.srvc.AddReservation(data).Send(w)

}

func (c *reservation_ctrl) UpdateReservation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	paymentCode, ok := vars["payment_code"]

	if !ok {
		libs.GetResponse("Payment Code not found in request", 400, true).Send(w)
		return
	}

	var data *models.Reservation

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		libs.GetResponse(err.Error(), 500, true).Send(w)
		return
	}

	c.srvc.UpdateReservation(data, paymentCode).Send(w)

}
