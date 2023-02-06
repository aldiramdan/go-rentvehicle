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

	c.srvc.GetAllHistories(user_id.(uint64)).Send(w)

}

func (c *history_ctrl) GetPageHistories(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user")

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

	c.srvc.GetPageHistories(user_id.(uint64), page, perpage).Send(w)

}

func (c *history_ctrl) GetHistoryById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true).Send(w)
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

	c.srvc.SearchHistory(user_id.(uint64), query[0]).Send(w)

}

func (c *history_ctrl) AddHistory(w http.ResponseWriter, r *http.Request) {

	var data *models.History

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		libs.GetResponse(err.Error(), 500, true).Send(w)
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

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	var data *models.History

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		libs.GetResponse(err.Error(), 500, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(data)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.srvc.UpdateHistory(data, id).Send(w)

}

func (c *history_ctrl) DeleteHistory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.srvc.DeleteHistory(id).Send(w)

}
