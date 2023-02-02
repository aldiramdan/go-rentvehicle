package histories

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RouteHistory(rt *mux.Router, db *gorm.DB) {

	route := rt.PathPrefix("/history").Subrouter()

	repo := NewRepo(db)
	srvc := NewSrvc(repo)
	ctrl := NewCtrl(srvc)

	route.HandleFunc("/", ctrl.GetAllHistories).Methods("GET")
	route.HandleFunc("/{id}", ctrl.GetHistoryById).Methods("GET")
	route.HandleFunc("/search/", ctrl.SearchHistory).Methods("GET")
	route.HandleFunc("/", ctrl.AddHistory).Methods("POST")
	route.HandleFunc("/{id}", ctrl.UpdateHistory).Methods("PUT")
	route.HandleFunc("/{id}", ctrl.DeleteHistory).Methods("DElETE")

}
