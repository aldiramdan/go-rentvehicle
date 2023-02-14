package histories

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

type history_ctrl struct {
	srvc interfaces.HistorySrvc
}

func NewCtrl(srvc interfaces.HistorySrvc) *history_ctrl {

	return &history_ctrl{srvc}

}

func (c *history_ctrl) GetAllHistories(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user")

	c.srvc.GetAllHistories(user_id.(string)).Send(w)

}

func (c *history_ctrl) GetPageHistories(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user")

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

	c.srvc.GetPageHistories(user_id.(string), page, perpage).Send(w)

}

func (c *history_ctrl) GetHistoryById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	c.srvc.GetHistoryById(id).Send(w)

}

func (c *history_ctrl) SearchHistory(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user")

	vars := r.URL.Query()

	query, ok := vars["s"]
	if !ok {
		libs.GetResponse("Missing query parameter", 400, true).Send(w)
		return
	}

	c.srvc.SearchHistory(user_id.(string), query[0]).Send(w)

}

func (c *history_ctrl) AddHistory(w http.ResponseWriter, r *http.Request) {

	var data *models.History

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

	c.srvc.AddHistory(data).Send(w)

}

func (c *history_ctrl) UpdateHistory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	var data *models.History

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

	c.srvc.UpdateHistory(data, id).Send(w)

}

func (c *history_ctrl) UpdateHistoryRating(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	var data *models.History

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

	c.srvc.UpdateHistoryRating(data, id).Send(w)

}

func (c *history_ctrl) DeleteHistory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	c.srvc.DeleteHistory(id).Send(w)

}
