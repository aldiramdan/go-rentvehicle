package categories

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

type category_ctrl struct {
	repo interfaces.CategoryRepo
}

func NewCategoryCtrl(repo interfaces.CategoryRepo) *category_ctrl {

	return &category_ctrl{repo}

}

func (c *category_ctrl) GetAllCategories(w http.ResponseWriter, r *http.Request) {

	result, err := c.repo.GetAllCategories()

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(*result) == 0 {
		res.ResponseError(w, http.StatusNotFound, "Data category empty")
		return
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *category_ctrl) GetCategoryById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := c.repo.GetCategoryById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.ResponseError(w, http.StatusNotFound, "Category not found")
			return
		default:
			res.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *category_ctrl) SearchCategories(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()

	query, ok := vars["s"]
	if !ok {
		res.ResponseError(w, http.StatusBadRequest, "Missing query parameter")
		return
	}

	result, err := c.repo.SearchCategories(query[0])

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(*result) == 0 {
		res.ResponseError(w, http.StatusNotFound, "Category not found")
		return
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *category_ctrl) AddCategory(w http.ResponseWriter, r *http.Request) {

	var data models.Category

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := c.repo.AddCategory(&data)

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res.ResponseJson(w, http.StatusCreated, result)

}

func (c *category_ctrl) UpdateCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var data models.Category

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = c.repo.GetCategoryById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.ResponseError(w, http.StatusNotFound, "Category not found")
			return
		default:
			res.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	result, err := c.repo.UpdateCategory(&data, id)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	res.ResponseJson(w, http.StatusOK, result)

}

func (c *category_ctrl) DeleteCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = c.repo.GetCategoryById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.ResponseError(w, http.StatusNotFound, "Category not found")
			return
		default:
			res.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	_, err = c.repo.DeleteCategory(id)

	if err != nil {
		res.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := map[string]string{"message": "Category deleted successfully"}

	res.ResponseJson(w, http.StatusOK, response)

}
