package interfaces

import (
	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/libs"
)

type VehicleRepo interface {
	GetAllVehicles() (*models.Vehicles, error)
	GetPageVehicles(limit, offset int) (*models.Vehicles, error)
	GetVehicleById(id string) (*models.Vehicle, error)
	GetPopulerVehicle() (*models.Vehicles, error)
	SearchVehicle(query string) (*models.Vehicles, error)
	AddVehicle(data *models.Vehicle) (*models.Vehicle, error)
	UpdateVehicle(data *models.Vehicle, id string) (*models.Vehicle, error)
	DeleteVehicle(id string) (*models.Vehicle, error)
}

type VehicleSrvc interface {
	GetAllVehicles() *libs.Response
	GetPageVehicles(page, perpage int) *libs.Response
	GetVehicleById(id string) *libs.Response
	GetPopulerVehicle() *libs.Response
	SearchVehicle(query string) *libs.Response
	AddVehicle(data *models.Vehicle) *libs.Response
	UpdateVehicle(data *models.Vehicle, id string) *libs.Response
	DeleteVehicle(id string) *libs.Response
}
