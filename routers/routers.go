package routers

import (
	"net/http"

	"github.com/aldiramdan/go-backend/databases/orm"
	"github.com/aldiramdan/go-backend/modules/v1/auth"
	"github.com/aldiramdan/go-backend/modules/v1/categories"
	"github.com/aldiramdan/go-backend/modules/v1/histories"
	"github.com/aldiramdan/go-backend/modules/v1/reservations"
	"github.com/aldiramdan/go-backend/modules/v1/users"
	"github.com/aldiramdan/go-backend/modules/v1/vehicles"
	"github.com/gorilla/mux"
)

func IndexRoute() (*mux.Router, error) {

	mainRoute := mux.NewRouter()

	db, err := orm.ConnectDB()

	if err != nil {
		return nil, err
	}

	var imgServer = http.FileServer(http.Dir("./public/"))
	mainRoute.PathPrefix("/public/").Handler(http.StripPrefix("/public/", imgServer))

	auth.RouteAuth(mainRoute, db)
	users.RouteUsers(mainRoute, db)
	vehicles.RouteVehicle(mainRoute, db)
	categories.RouteCategory(mainRoute, db)
	histories.RouteHistory(mainRoute, db)
	reservations.RouteTransaction(mainRoute, db)

	return mainRoute, nil

}
