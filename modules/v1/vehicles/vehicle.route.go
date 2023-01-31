package vehicles

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RouteVehicle(rt *mux.Router, db *gorm.DB) {

	route := rt.PathPrefix("/vehicle").Subrouter()

	repo := NewVehicleRepo(db)
	ctrl := NewVehicleCtrl(repo)

	route.HandleFunc("/", ctrl.GetAllVehicles).Methods("GET")
	route.HandleFunc("/popular/", ctrl.GetPopulerVehicle).Methods("GET")
	route.HandleFunc("/{id}", ctrl.GetVehicleById).Methods("GET")
	route.HandleFunc("/search/", ctrl.SearchVehicle).Methods("GET")
	route.HandleFunc("/", ctrl.AddVehicle).Methods("POST")
	route.HandleFunc("/{id}", ctrl.UpdateVehicle).Methods("PUT")
	route.HandleFunc("/{id}", ctrl.DeleteVehicle).Methods("DElETE")

}
