package vehicles

import (
	"errors"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"gorm.io/gorm"
)

type vehicle_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *vehicle_repo {
	return &vehicle_repo{db}
}

func (r *vehicle_repo) GetAllVehicles() (*models.Vehicles, error) {

	var data models.Vehicles

	if err := r.db.
		Preload("Category").
		Order("created_at DESC").
		Find(&data).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	if len(data) == 0 {
		return nil, errors.New("data vehicle is empty")
	}

	return &data, nil

}

func (r *vehicle_repo) GetPageVehicles(limit, offset int) (*models.Vehicles, error) {

	var data models.Vehicles

	if err := r.db.
		Preload("Category").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&data).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	if len(data) == 0 {
		return nil, errors.New("data vehicle is empty")
	}

	return &data, nil

}

func (r *vehicle_repo) GetVehicleById(id string) (*models.Vehicle, error) {

	var data models.Vehicle

	if err := r.db.
		Preload("Category").
		First(&data, "vehicle_id = ?", id).Error; err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *vehicle_repo) GetPopulerVehicle() (*models.Vehicles, error) {

	var data models.Vehicles

	if err := r.db.
		Preload("Category").
		Order("rating DESC, total_rent DESC").
		Limit(5).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *vehicle_repo) SearchVehicle(query string) (*models.Vehicles, error) {

	var data models.Vehicles

	if err := r.db.
		Preload("Category").
		Order("created_at DESC").
		Joins("JOIN category ON category.category_id = vehicle.category_id").
		Where("LOWER(vehicle.name) LIKE ? OR LOWER(location) LIKE ? OR LOWER(category.name) LIKE ?", "%"+query+"%", "%"+query+"%", ""+query+"").
		Find(&data).Error; err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("search data vehicle not found")
	}

	return &data, nil

}

func (r *vehicle_repo) AddVehicle(data *models.Vehicle) (*models.Vehicle, error) {

	var dataCategory models.Category
	if err := r.db.
		First(&dataCategory, "category_id = ?", data.CategoryID).Error; err != nil {
		return nil, errors.New("data category not found")
	}

	if err := r.db.
		Preload("Category").
		Create(data).
		Find(&data).Error; err != nil {
		return nil, errors.New("failed to create data")
	}

	return data, nil

}

func (r *vehicle_repo) UpdateVehicle(data *models.Vehicle, id string) (*models.Vehicle, error) {

	var dataCategory models.Category
	if err := r.db.
		First(&dataCategory, "category_id", data.CategoryID).Error; err != nil {
		return nil, errors.New("data category not found")
	}

	if err := r.db.
		Model(&data).
		Preload("Category").
		Where("vehicle_id = ?", id).
		Updates(&data).
		Find(&data).Error; err != nil {
		return nil, errors.New("failed to update data")
	}

	return data, nil

}

func (r *vehicle_repo) DeleteVehicle(id string) (*models.Vehicle, error) {

	var data models.Vehicle

	if err := r.db.
		Delete(data, "vehicle_id = ?", id).Error; err != nil {
		return nil, errors.New("failed to delete data")
	}

	return &data, nil

}
