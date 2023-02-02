package reservations

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RouteTransaction(rt *mux.Router, db *gorm.DB) {

	route := rt.PathPrefix("/reservation").Subrouter()

	repo := NewRepo(db)
	srvc := NewSrvc(repo)
	ctrl := NewCtrl(srvc)

	route.HandleFunc("/", ctrl.GetAllReservations).Methods("GET")
	route.HandleFunc("/{id}", ctrl.GetReservationById).Methods("GET")
	route.HandleFunc("/", ctrl.AddReservation).Methods("POST")
	route.HandleFunc("/payment/{id}", ctrl.Payment).Methods("PUT")

}
