package reservations

import (
	"errors"

	"github.com/aldiramdan/go-backend/databases/orm/models"
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
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, email, name, phone, created_at, updated_at")
		}).
		Preload("Vehicle", func(db *gorm.DB) *gorm.DB {
			return db.Select("vehicle_id, name, location, price, category_id, rating, created_at, updated_at")
		}).
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

func (r *reservation_repo) GetReservationUser(user_id string) (*models.Reservation, error) {

	var data models.Reservation

	if err := r.db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, email, name, phone, created_at, updated_at")
		}).
		Preload("Vehicle", func(db *gorm.DB) *gorm.DB {
			return db.Select("vehicle_id, name, location, price, category_id, rating, created_at, updated_at")
		}).
		Preload("Vehicle.Category").
		Where("user_id = ? AND payment_status = ?", user_id, "Pending").
		Find(&data).Error; err != nil {
		return nil, err
	}

	if data.PaymentStatus == "" || data.UserID == "" {
		return nil, errors.New("data reservation is empty")
	}

	return &data, nil

}

func (r *reservation_repo) GetPageReservations(limit, offset int) (*models.Reservations, error) {

	var data models.Reservations

	if err := r.db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, email, name, phone, created_at, updated_at")
		}).
		Preload("Vehicle", func(db *gorm.DB) *gorm.DB {
			return db.Select("vehicle_id, name, location, price, category_id, rating, created_at, updated_at")
		}).
		Preload("Vehicle.Category").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&data).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	if len(data) == 0 {
		return nil, errors.New("data reservation is empty")
	}

	return &data, nil

}

func (r *reservation_repo) GetReservationById(id string) (*models.Reservation, error) {

	var data models.Reservation

	if err := r.db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, email, name, phone, created_at, updated_at")
		}).
		Preload("Vehicle", func(db *gorm.DB) *gorm.DB {
			return db.Select("vehicle_id, name, location, price, category_id, rating, created_at, updated_at")
		}).
		Preload("Vehicle.Category").
		First(&data, "reservation_id = ?", id).Error; err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *reservation_repo) GetReservationByCode(paymentCode string) (*models.Reservation, error) {

	var data models.Reservation

	if err := r.db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, email, name, phone, created_at, updated_at")
		}).
		Preload("Vehicle", func(db *gorm.DB) *gorm.DB {
			return db.Select("vehicle_id, name, location, price, category_id, rating, created_at, updated_at")
		}).
		Preload("Vehicle.Category").
		First(&data, "payment_code = ?", paymentCode).Error; err != nil {
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
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, email, name, phone, created_at, updated_at")
		}).
		Preload("Vehicle", func(db *gorm.DB) *gorm.DB {
			return db.Select("vehicle_id, name, location, price, category_id, rating, created_at, updated_at")
		}).
		Preload("Vehicle.Category").
		First(&data).Error; err != nil {
		return nil, err
	}

	return data, tx.Commit().Error

}

func (r *reservation_repo) BeforeCreate(data *models.Reservation) error {

	var dataUser models.User
	if err := r.db.
		First(&dataUser, "user_id = ?", data.UserID).Error; err != nil {
		return errors.New("data user not found")
	}

	var dataVehicle models.Vehicle
	if err := r.db.
		First(&dataVehicle, "vehicle_id = ?", data.VehicleID).Error; err != nil {
		return errors.New("data vehicle not found")
	}

	if dataVehicle.Status != "Available" {
		return errors.New("vehicle not available")
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
		First(&dataVehicle, "vehicle_id = ?", data.VehicleID).Error; err != nil {
		return errors.New("data vehicle not found")
	}

	if data.PaymentMethod == "Cash" {
		if err := r.db.
			Model(&dataVehicle).
			Where("vehicle_id = ?", data.VehicleID).
			Update("stock", dataVehicle.Stock-data.Quantity).Error; err != nil {
			return errors.New("failed update data vehicle")
		}

	}

	if err := r.db.
		Model(&dataVehicle).
		Where("vehicle_id = ?", data.VehicleID).
		Update("total_rent", dataVehicle.TotalRent+1).Error; err != nil {
		return errors.New("failed update data vehicle")
	}

	return nil

}

func (r *reservation_repo) UpdateReservation(data *models.Reservation, paymentCode string) (*models.Reservation, error) {

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

	if err := r.BeforeUpdate(data, paymentCode); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Save(data).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := r.AfterUpdate(data, paymentCode); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, email, name, phone, created_at, updated_at")
		}).
		Preload("Vehicle", func(db *gorm.DB) *gorm.DB {
			return db.Select("vehicle_id, name, location, price, category_id, rating, created_at, updated_at")
		}).
		Preload("Vehicle.Category").
		First(&data).Error; err != nil {
		return nil, err
	}

	return data, tx.Commit().Error

}

func (r *reservation_repo) BeforeUpdate(data *models.Reservation, paymentCode string) error {

	if err := r.db.
		Model(data).
		Where("payment_code = ?", paymentCode).
		Updates(&data).Error; err != nil {
		return errors.New("failed to update data")
	}

	if err := r.db.
		First(&data, "payment_code = ?", paymentCode).Error; err != nil {
		return errors.New("data transaction not found")
	}

	return nil

}

func (r *reservation_repo) AfterUpdate(data *models.Reservation, paymentCode string) error {

	var dataVehicle models.Vehicle

	if err := r.db.
		First(&dataVehicle, "vehicle_id = ?", data.VehicleID).Error; err != nil {
		return errors.New("data vehicle not found")
	}

	if dataVehicle.Stock == 0 {
		if err := r.db.
			Model(&dataVehicle).
			Where("vehicle_id = ?", data.VehicleID).
			Update("status", "unavailable").Error; err != nil {
			return errors.New("failed update data vehicle")
		}
	}

	if err := r.db.
		Model(&dataVehicle).
		Where("vehicle_id = ?", data.VehicleID).
		Update("stock", dataVehicle.Stock-data.Quantity).Error; err != nil {
		return errors.New("failed update data vehicle")
	}

	return nil

}
