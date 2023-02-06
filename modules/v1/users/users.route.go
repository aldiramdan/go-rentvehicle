package users

import (
	"github.com/aldiramdan/go-backend/middlewares"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RouteUsers(rt *mux.Router, db *gorm.DB) {

	route := rt.PathPrefix("/users").Subrouter()

	repo := NewRepo(db)
	srvc := NewSrvc(repo)
	ctrl := NewCtrl(srvc)

	route.HandleFunc("/", middlewares.Handle(ctrl.AddUser, middlewares.AuthUploadFile())).Methods("POST")
	route.HandleFunc("/p", middlewares.Handle(ctrl.GetAllUsers, middlewares.AuthMidle("admin"))).Methods("GET")
	route.HandleFunc("/", middlewares.Handle(ctrl.GetAllUsers, middlewares.AuthMidle("admin"))).Methods("GET")
	route.HandleFunc("/profile", middlewares.Handle(ctrl.GetUserById, middlewares.AuthMidle("user", "admin"))).Methods("GET")
	route.HandleFunc("/profile/update", middlewares.Handle(ctrl.UpdateUser, middlewares.AuthUploadFile(), middlewares.AuthMidle("user", "admin"))).Methods("PUT")
	route.HandleFunc("/profile/delete", middlewares.Handle(ctrl.DeleteUser, middlewares.AuthMidle("user", "admin"))).Methods("DElETE")

}
