package categories

import (
	"github.com/aldiramdan/go-backend/middlewares"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RouteCategory(rt *mux.Router, db *gorm.DB) {

	route := rt.PathPrefix("/vehicle/category").Subrouter()

	repo := NewRepo(db)
	srvc := NewSrvc(repo)
	ctrl := NewCtrl(srvc)

	route.HandleFunc("/", ctrl.GetAllCategories).Methods("GET")
	route.HandleFunc("/", middlewares.Handle(ctrl.UpdateCategory, middlewares.AuthMidle("admin"))).Methods("POST")
	route.HandleFunc("/search/", ctrl.SearchCategories).Methods("GET")
	route.HandleFunc("/{id}", ctrl.GetCategoryById).Methods("GET")
	route.HandleFunc("/{id}", middlewares.Handle(ctrl.UpdateCategory, middlewares.AuthMidle("admin"))).Methods("PUT")
	route.HandleFunc("/{id}", middlewares.Handle(ctrl.DeleteCategory, middlewares.AuthMidle("admin"))).Methods("DElETE")

}
