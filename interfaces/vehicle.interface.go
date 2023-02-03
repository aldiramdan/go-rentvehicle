package interfaces

import (
	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/libs"
)

type VehicleRepo interface {
	GetAllVehicles() (*models.Vehicles, error)
	GetVehicleById(id uint64) (*models.Vehicle, error)
	GetPopulerVehicle() (*models.Vehicles, error)
	SearchVehicle(query string) (*models.Vehicles, error)
	AddVehicle(data *models.Vehicle) (*models.Vehicle, error)
	UpdateVehicle(data *models.Vehicle, id uint64) (*models.Vehicle, error)
	DeleteVehicle(id uint64) (*models.Vehicle, error)
}

type VehicleSrvc interface {
	GetAllVehicles() *libs.Response
	GetVehicleById(id uint64) *libs.Response
	GetPopulerVehicle() *libs.Response
	SearchVehicle(query string) *libs.Response
	AddVehicle(data *models.Vehicle) *libs.Response
	UpdateVehicle(data *models.Vehicle, id uint64) *libs.Response
	DeleteVehicle(id uint64) *libs.Response
}
