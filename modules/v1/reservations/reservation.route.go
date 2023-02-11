package reservations

import (
	"github.com/aldiramdan/go-backend/middlewares"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RouteTransaction(rt *mux.Router, db *gorm.DB) {

	route := rt.PathPrefix("/reservation").Subrouter()

	repo := NewRepo(db)
	srvc := NewSrvc(repo)
	ctrl := NewCtrl(srvc)

	route.HandleFunc("/", middlewares.Handle(ctrl.GetAllReservations, middlewares.AuthMidle("admin"))).Methods("GET")
	route.HandleFunc("/p", middlewares.Handle(ctrl.GetPageReservations, middlewares.AuthMidle("admin"))).Methods("GET")
	route.HandleFunc("/id={id}", middlewares.Handle(ctrl.GetReservationById, middlewares.AuthMidle("admin"))).Methods("GET")
	route.HandleFunc("/{payment_code}/", middlewares.Handle(ctrl.GetReservationByCode, middlewares.AuthMidle("user"))).Methods("GET")
	route.HandleFunc("/", middlewares.Handle(ctrl.AddReservation, middlewares.AuthMidle("user"))).Methods("POST")
	route.HandleFunc("/callback/{payment_code}", middlewares.Handle(ctrl.UpdateReservation, middlewares.AuthMidle("user", "admin"))).Methods("PUT")

}
