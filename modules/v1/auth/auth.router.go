package auth

import (
	"github.com/aldiramdan/go-backend/modules/v1/users"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RouteAuth(rt *mux.Router, db *gorm.DB) {

	route := rt.PathPrefix("/auth").Subrouter()

	repo := users.NewRepo(db)
	srvc := NewSrvc(repo)
	ctrl := NewCtrl(srvc)

	route.HandleFunc("/login", ctrl.Login).Methods("POST")
	route.HandleFunc("/confirm_email/{token}", ctrl.VerifyEmail).Methods("GET")
	route.HandleFunc("/resend_email/", ctrl.ResendEmail).Methods("POST")

}
