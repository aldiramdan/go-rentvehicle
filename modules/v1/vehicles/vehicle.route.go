package vehicles

import (
	"github.com/aldiramdan/go-backend/middlewares"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RouteVehicle(rt *mux.Router, db *gorm.DB) {

	route := rt.PathPrefix("/vehicle").Subrouter()

	repo := NewRepo(db)
	srvc := NewSrvc(repo)
	ctrl := NewCtrl(srvc)

	route.HandleFunc("/", ctrl.GetAllVehicles).Methods("GET")
	route.HandleFunc("/", middlewares.Handle(ctrl.AddVehicle, middlewares.AuthMidle("admin"))).Methods("POST")
	route.HandleFunc("/search/", ctrl.SearchVehicle).Methods("GET")
	route.HandleFunc("/popular", ctrl.GetPopulerVehicle).Methods("GET")
	route.HandleFunc("/{id}", ctrl.GetVehicleById).Methods("GET")
	route.HandleFunc("/{id}", middlewares.Handle(ctrl.UpdateVehicle, middlewares.AuthMidle("admin"))).Methods("PUT")
	route.HandleFunc("/{id}", middlewares.Handle(ctrl.DeleteVehicle, middlewares.AuthMidle("admin"))).Methods("DElETE")

}
