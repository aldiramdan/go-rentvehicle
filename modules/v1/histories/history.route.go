package histories

import (
	"github.com/aldiramdan/go-backend/middlewares"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RouteHistory(rt *mux.Router, db *gorm.DB) {

	route := rt.PathPrefix("/history").Subrouter()

	repo := NewRepo(db)
	srvc := NewSrvc(repo)
	ctrl := NewCtrl(srvc)

	route.HandleFunc("/", middlewares.Handle(ctrl.GetAllHistories, middlewares.AuthMidle("user", "admin"))).Methods("GET")
	route.HandleFunc("/p", middlewares.Handle(ctrl.GetPageHistories, middlewares.AuthMidle("user", "admin"))).Methods("GET")
	route.HandleFunc("/{id}", middlewares.Handle(ctrl.GetHistoryById, middlewares.AuthMidle("admin"))).Methods("GET")
	route.HandleFunc("/search/", middlewares.Handle(ctrl.SearchHistory, middlewares.AuthMidle("user", "admin"))).Methods("GET")
	route.HandleFunc("/", middlewares.Handle(ctrl.AddHistory, middlewares.AuthMidle("admin"))).Methods("POST")
	route.HandleFunc("/{id}", middlewares.Handle(ctrl.UpdateHistory, middlewares.AuthMidle("admin"))).Methods("PUT")
	route.HandleFunc("/{id}", middlewares.Handle(ctrl.DeleteHistory, middlewares.AuthMidle("admin"))).Methods("DElETE")

}
