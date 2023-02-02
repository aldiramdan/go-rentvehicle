package histories

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/interfaces"
	res "github.com/aldiramdan/go-backend/libs"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type history_ctrl struct {
	repo interfaces.HistoryRepo
}

func NewHistoryCtrl(repo interfaces.HistoryRepo) *history_ctrl {

	return &history_ctrl{repo}

}

func (c *history_ctrl) GetAllHistories(w http.ResponseWriter, r *http.Request) {

	result, err := c.repo.GetAllHistories()

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if len(*result) == 0 {
		res.ResponseError(w, http.StatusNotFound, "Data history empty")
		return
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *history_ctrl) GetHistoryById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := c.repo.GetHistoryById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.ResponseError(w, http.StatusNotFound, "Hisotry not found")
			return
		default:
			res.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *history_ctrl) SearchHistory(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()

	query, ok := vars["s"]
	if !ok {
		res.ResponseError(w, http.StatusBadRequest, "Missing query parameter")
		return
	}

	result, err := c.repo.SearchHistory(query[0])

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(*result) == 0 {
		res.ResponseError(w, http.StatusNotFound, "Data not found")
		return
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *history_ctrl) AddHistory(w http.ResponseWriter, r *http.Request) {

	var data models.History

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := c.repo.AddHistory(&data)
	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *history_ctrl) UpdateHistory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var data models.History

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = c.repo.GetHistoryById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.ResponseError(w, http.StatusNotFound, "Hisotry not found")
			return
		default:
			res.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	result, err := c.repo.UpdateHistory(&data, id)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *history_ctrl) DeleteHistory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = c.repo.GetHistoryById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.ResponseError(w, http.StatusNotFound, "Hisotry not found")
			return
		default:
			res.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	_, err = c.repo.DeleteHistory(id)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := map[string]string{"message": "History deleted successfully"}

	res.ResponseJson(w, http.StatusOK, response)

}
