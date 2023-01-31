package categories

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RouteCategory(rt *mux.Router, db *gorm.DB) {

	route := rt.PathPrefix("/vehicle/category").Subrouter()

	repo := NewCategoryRepo(db)
	ctrl := NewCategoryCtrl(repo)

	route.HandleFunc("/", ctrl.GetAllCategories).Methods("GET")
	route.HandleFunc("/{id}", ctrl.GetCategoryById).Methods("GET")
	route.HandleFunc("/search/", ctrl.SearchCategories).Methods("GET")
	route.HandleFunc("/", ctrl.AddCategory).Methods("POST")
	route.HandleFunc("/{id}", ctrl.UpdateCategory).Methods("PUT")
	route.HandleFunc("/{id}", ctrl.DeleteCategory).Methods("DElETE")

}
