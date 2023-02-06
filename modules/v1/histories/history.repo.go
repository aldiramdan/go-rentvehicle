package histories

import (
	"errors"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"gorm.io/gorm"
)

type history_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *history_repo {

	return &history_repo{db}

}

type vehicle_repo struct {
	db *gorm.DB
}

func NewVehicleRepo(db *gorm.DB) *vehicle_repo {
	return &vehicle_repo{db}
}

func (r *history_repo) GetAllHistories(user_id uint64) (*models.Histories, error) {

	var data models.Histories

	if err := r.db.
		Preload("Reservation").
		Preload("Reservation.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, email, name, phone")
		}).
		Preload("Reservation.Vehicle", func(db *gorm.DB) *gorm.DB {
			return db.Select("vehicle_id, name, location, price, category_id, rating")
		}).
		Preload("Reservation.Vehicle.Category").
		Order("created_at DESC").
		Joins("JOIN reservation ON reservation.reservation_id = history.reservation_id").
		Joins("JOIN users ON users.user_id = reservation.user_id").
		Where("reservation.user_id = ?", user_id).
		Find(&data).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	if len(data) == 0 {
		return nil, errors.New("data history is empty")
	}

	return &data, nil

}

func (r *history_repo) GetPageHistories(user_id uint64, limit, offset int) (*models.Histories, error) {

	var data models.Histories

	if err := r.db.
		Preload("Reservation").
		Preload("Reservation.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, email, name, phone")
		}).
		Preload("Reservation.Vehicle", func(db *gorm.DB) *gorm.DB {
			return db.Select("vehicle_id, name, location, price, category_id, rating")
		}).
		Preload("Reservation.Vehicle.Category").
		Order("created_at DESC").
		Joins("JOIN reservation ON reservation.reservation_id = history.reservation_id").
		Joins("JOIN users ON users.user_id = reservation.user_id").
		Where("reservation.user_id = ?", user_id).
		Limit(limit).
		Offset(offset).
		Find(&data).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	if len(data) == 0 {
		return nil, errors.New("data history is empty")
	}

	return &data, nil

}

func (r *history_repo) GetHistoryById(id uint64) (*models.History, error) {

	var data models.History

	if err := r.db.
		Preload("Reservation").
		Preload("Reservation.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, email, name, phone")
		}).
		Preload("Reservation.Vehicle", func(db *gorm.DB) *gorm.DB {
			return db.Select("vehicle_id, name, location, price, category_id, rating")
		}).
		Preload("Reservation.Vehicle.Category").
		First(&data, id).Error; err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *history_repo) SearchHistory(user_id uint64, query string) (*models.Histories, error) {

	var data models.Histories

	if err := r.db.
		Preload("Reservation").
		Preload("Reservation.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, email, name, phone")
		}).
		Preload("Reservation.Vehicle", func(db *gorm.DB) *gorm.DB {
			return db.Select("vehicle_id, name, location, price, category_id, rating")
		}).
		Preload("Reservation.Vehicle.Category").
		Order("created_at DESC").
		Joins("JOIN reservation ON reservation.reservation_id = history.reservation_id").
		Joins("JOIN vehicle ON vehicle.vehicle_id = reservation.vehicle_id").
		Joins("JOIN users ON users.user_id = reservation.user_id").
		Where("LOWER(vehicle.name) LIKE ? AND reservation.user_id = ?", "%"+query+"%", user_id).
		Find(&data).Error; err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("search data history not found")
	}

	return &data, nil

}

func (r *history_repo) AddHistory(data *models.History) (*models.History, error) {

	var dataReservation models.Reservation
	if err := r.db.
		First(&dataReservation, data.ReservationID).Error; err != nil {
		return nil, errors.New("data reservation not found")
	}

	if err := r.db.
		Preload("Reservation").
		Preload("Reservation.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, email, name, phone")
		}).
		Preload("Reservation.Vehicle", func(db *gorm.DB) *gorm.DB {
			return db.Select("vehicle_id, name, location, price, category_id, rating")
		}).
		Preload("Reservation.Vehicle.Category").
		Create(data).
		Find(&data).Error; err != nil {
		return nil, errors.New("failed to create data")
	}

	return data, nil

}

func (r *history_repo) UpdateHistory(data *models.History, id uint64) (*models.History, error) {

	var dataReservation models.Reservation
	if err := r.db.
		First(&dataReservation, id).Error; err != nil {
		return nil, errors.New("data reservation not found")
	}

	if err := r.db.
		Model(&data).
		Preload("Reservation").
		Preload("Reservation.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, email, name, phone")
		}).
		Preload("Reservation.Vehicle", func(db *gorm.DB) *gorm.DB {
			return db.Select("vehicle_id, name, location, price, category_id, rating")
		}).
		Preload("Reservation.Vehicle.Category").
		Where("history_id = ?", id).
		Updates(&data).
		Find(&data).Error; err != nil {
		return nil, errors.New("failed to update data")
	}

	return data, nil

}

func (r *history_repo) DeleteHistory(id uint64) (*models.History, error) {

	var data models.History

	if err := r.db.
		Delete(data, id).Error; err != nil {
		return nil, err
	}

	return &data, nil

}
