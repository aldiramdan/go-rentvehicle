package histories

import (
	"errors"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"gorm.io/gorm"
)

type history_repo struct {
	db *gorm.DB
}

func NewHisotryRepo(db *gorm.DB) *history_repo {

	return &history_repo{db}

}

type vehicle_repo struct {
	db *gorm.DB
}

func NewVehicleRepo(db *gorm.DB) *vehicle_repo {
	return &vehicle_repo{db}
}

func (r *history_repo) GetAllHistories() (*models.Histories, error) {

	var data models.Histories

	if err := r.db.Preload("User").Preload("Vehicle").Preload("Vehicle.Category").Order("created_at DESC").Find(&data).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	return &data, nil

}

func (r *history_repo) GetHistoryById(id uint64) (*models.History, error) {

	var data models.History

	if err := r.db.Preload("User").Preload("Vehicle").Preload("Vehicle.Category").First(&data, id).Error; err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *history_repo) SearchHistory(query string) (*models.Histories, error) {

	var data models.Histories

	if err := r.db.Preload("User").Preload("Vehicle").Preload("Vehicle.Category").Order("created_at DESC").Joins("JOIN vehicle ON vehicle.vehicle_id = history.vehicle_id AND LOWER(vehicle.name) LIKE ?", "%"+query+"%").Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *history_repo) AddHistory(data *models.History) (*models.History, error) {

	var dataUser models.User
	if err := r.db.First(&dataUser, data.UserID).Error; err != nil {
		return nil, errors.New("data user not found")
	}

	var dataVehicle models.Vehicle
	if err := r.db.First(&dataVehicle, data.VehicleID).Error; err != nil {
		return nil, errors.New("data vehicle not found")
	}

	if err := r.db.Preload("User").Preload("Vehicle").Preload("Vehicle.Category").Create(data).Find(&data).Error; err != nil {
		return nil, errors.New("failed to create data")
	}

	return data, nil

}

func (r *history_repo) UpdateHistory(data *models.History, id uint64) (*models.History, error) {

	var dataUser models.User
	if err := r.db.First(&dataUser, data.UserID).Error; err != nil {
		return nil, errors.New("data user not found")
	}

	var dataVehicle models.Vehicle
	if err := r.db.First(&dataVehicle, data.VehicleID).Error; err != nil {
		return nil, errors.New("data vehicle not found")
	}

	if err := r.db.Model(&data).Preload("User").Preload("Vehicle").Preload("Vehicle.Category").Where("history_id = ?", id).Updates(&data).Find(&data).Error; err != nil {
		return nil, errors.New("failed to update data")
	}

	return data, nil

}

func (r *history_repo) DeleteHistory(id uint64) (*models.History, error) {

	var data models.History

	if err := r.db.Delete(data, id).Error; err != nil {
		return nil, err
	}

	return &data, nil

}
