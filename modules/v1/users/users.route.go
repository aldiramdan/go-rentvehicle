package users

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RouteUsers(rt *mux.Router, db *gorm.DB) {

	route := rt.PathPrefix("/users").Subrouter()

	repo := NewRepo(db)
	srvc := NewSrvc(repo)
	ctrl := NewCtrl(srvc)

	route.HandleFunc("/", ctrl.GetAllUsers).Methods("GET")
	route.HandleFunc("/{id}", ctrl.GetUserById).Methods("GET")
	route.HandleFunc("/", ctrl.AddUser).Methods("POST")
	route.HandleFunc("/{id}", ctrl.UpdateUser).Methods("PUT")
	route.HandleFunc("/{id}", ctrl.DeleteUser).Methods("DElETE")

}
