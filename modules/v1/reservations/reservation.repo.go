package reservations

import (
	"errors"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/libs"
	"gorm.io/gorm"
)

type reservation_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *reservation_repo {

	return &reservation_repo{db}

}

func (r *reservation_repo) GetAllReservations() (*models.Reservations, error) {

	var data models.Reservations

	if err := r.db.
		Preload("User").
		Preload("Vehicle").
		Preload("Vehicle.Category").
		Order("created_at DESC").
		Find(&data).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	if len(data) == 0 {
		return nil, errors.New("data reservation is empty")
	}

	return &data, nil

}

func (r *reservation_repo) GetReservationById(id uint64) (*models.Reservation, error) {

	var data models.Reservation

	if err := r.db.
		Preload("User").
		Preload("Vehicle").
		Preload("Vehicle.Category").
		First(&data, id).Error; err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *reservation_repo) AddReservation(data *models.Reservation) (*models.Reservation, error) {

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := r.BeforeCreate(data); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(data).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	history := &models.History{
		ReservationID: data.ReservationID,
	}

	if err := tx.Create(history).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to create history")
	}

	if err := r.AfterCreate(data); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.
		Preload("User").
		Preload("Vehicle").
		Preload("Vehicle.Category").
		First(&data).Error; err != nil {
		return nil, err
	}

	return data, tx.Commit().Error

}

func (r *reservation_repo) BeforeCreate(data *models.Reservation) error {

	var dataUser models.User
	if err := r.db.
		First(&dataUser, data.UserID).Error; err != nil {
		return errors.New("data user not found")
	}

	var dataVehicle models.Vehicle
	if err := r.db.
		First(&dataVehicle, data.VehicleID).Error; err != nil {
		return errors.New("data vehicle not found")
	}

	qty := dataVehicle.Stock - data.Quantity

	if int(qty) < 0 {
		return errors.New("stock not enough")
	}

	return nil

}

func (r *reservation_repo) AfterCreate(data *models.Reservation) error {

	var dataVehicle models.Vehicle

	if err := r.db.
		First(&dataVehicle, data.VehicleID).Error; err != nil {
		return errors.New("data vehicle not found")
	}

	newRating := libs.CalculateNewRating(dataVehicle.TotalRent, dataVehicle.Rating, data.Rating)

	if err := r.db.
		Model(&dataVehicle).
		Where("vehicle_id = ?", data.VehicleID).
		Updates(map[string]interface{}{"total_rent": dataVehicle.TotalRent + 1, "rating": newRating}).Error; err != nil {
		return errors.New("failed update data vehicle")
	}

	return nil

}

func (r *reservation_repo) Payment(data *models.Reservation, id uint64) (*models.Reservation, error) {

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := r.BeforeUpdate(data, id); err != nil {
		tx.Rollback()
		return nil, err
	}

	data.PaymentStatus = "Paid"
	data.IsBooked = true

	if err := tx.Save(data).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	history := &models.History{
		ReservationID: data.ReservationID,
	}

	if err := tx.Save(history).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to create history")
	}

	if err := r.AfterUpdate(data); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.
		Preload("User").
		Preload("Vehicle").
		Preload("Vehicle.Category").
		First(&data).Error; err != nil {
		return nil, err
	}

	return data, tx.Commit().Error

}

func (r *reservation_repo) BeforeUpdate(data *models.Reservation, id uint64) error {

	if err := r.db.
		Model(data).
		Where("reservation_id = ?", id).
		Updates(&data).Error; err != nil {
		return errors.New("failed to update data")
	}

	if err := r.db.
		First(&data, id).Error; err != nil {
		return errors.New("data transaction not found")
	}

	if data.PaymentStatus == "Paid" {
		return errors.New("payment complated")
	}

	return nil

}

func (r *reservation_repo) AfterUpdate(data *models.Reservation) error {

	var dataVehicle models.Vehicle

	if err := r.db.
		First(&dataVehicle, data.VehicleID).Error; err != nil {
		return errors.New("data vehicle not found")
	}

	if err := r.db.
		Model(&dataVehicle).
		Where("vehicle_id = ?", data.VehicleID).
		Update("stock", dataVehicle.Stock-data.Quantity).Error; err != nil {
		return errors.New("failed update data vehicle")
	}

	if dataVehicle.Stock == 0 {
		if err := r.db.
			Model(&dataVehicle).
			Where("vehicle_id = ?", data.VehicleID).
			Updates(map[string]interface{}{"status": "unavailable"}).Error; err != nil {
			return errors.New("failed update data vehicle")
		}
	}

	return nil

}
