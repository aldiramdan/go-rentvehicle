package categories

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

type category_ctrl struct {
	srvc interfaces.CategorySrvc
}

func NewCtrl(srvc interfaces.CategorySrvc) *category_ctrl {

	return &category_ctrl{srvc}

}

func (c *category_ctrl) GetAllCategories(w http.ResponseWriter, r *http.Request) {

	c.srvc.GetAllCategories().Send(w)

}

func (c *category_ctrl) GetPageCategories(w http.ResponseWriter, r *http.Request) {

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

	c.srvc.GetPageCategories(page, perpage).Send(w)

}

func (c *category_ctrl) GetCategoryById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]
	
	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	c.srvc.GetCategoryById(id).Send(w)

}

func (c *category_ctrl) SearchCategories(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()

	query, ok := vars["s"]
	if !ok {
		libs.GetResponse("Missing query parameter", 400, true).Send(w)
		return
	}

	c.srvc.SearchCategories(query[0]).Send(w)

}

func (c *category_ctrl) AddCategory(w http.ResponseWriter, r *http.Request) {

	var data *models.Category

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

	c.srvc.AddCategory(data).Send(w)

}

func (c *category_ctrl) UpdateCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	var data *models.Category

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

	c.srvc.UpdateCategory(data, id).Send(w)

}

func (c *category_ctrl) DeleteCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	c.srvc.DeleteCategory(id).Send(w)

}
