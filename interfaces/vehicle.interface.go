package interfaces

import "github.com/aldiramdan/go-backend/databases/orm/models"

type VehicleRepo interface {
	GetAllVehicles() (*models.Vehicles, error)
	GetVehicleById(id uint64) (*models.Vehicle, error)
	GetPopulerVehicle() (*models.Vehicles, error)
	SearchVehicle(query string) (*models.Vehicles, error)
	AddVehicle(data *models.Vehicle) (*models.Vehicle, error)
	UpdateVehicle(data *models.Vehicle, id uint64) (*models.Vehicle, error)
	DeleteVehicle(id uint64) (*models.Vehicle, error)
}
