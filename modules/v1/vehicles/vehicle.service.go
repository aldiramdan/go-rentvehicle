package vehicles

import (
	"os"

	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/interfaces"
	"github.com/aldiramdan/go-backend/libs"
	"gorm.io/gorm"
)

type vehicle_service struct {
	repo interfaces.VehicleRepo
}

func NewSrvc(repo interfaces.VehicleRepo) *vehicle_service {

	return &vehicle_service{repo}

}

func (s *vehicle_service) GetAllVehicles() *libs.Response {

	result, err := s.repo.GetAllVehicles()

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *vehicle_service) GetVehicleById(id uint64) *libs.Response {

	result, err := s.repo.GetVehicleById(id)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *vehicle_service) GetPopulerVehicle() *libs.Response {

	result, err := s.repo.GetPopulerVehicle()

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *vehicle_service) SearchVehicle(query string) *libs.Response {

	result, err := s.repo.SearchVehicle(query)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}
func (s *vehicle_service) AddVehicle(data *models.Vehicle) *libs.Response {

	result, err := s.repo.AddVehicle(data)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}
func (s *vehicle_service) UpdateVehicle(data *models.Vehicle, id uint64) *libs.Response {

	datas, err := s.repo.GetVehicleById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	if datas.Picture != "public/default_image.jpg" {
		_ = os.Remove(datas.Picture)
	}

	result, err := s.repo.UpdateVehicle(data, id)

	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	return libs.GetResponse(result, 200, false)

}
func (s *vehicle_service) DeleteVehicle(id uint64) *libs.Response {

	data, err := s.repo.GetVehicleById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	if data.Picture != "public/default_image.jpg" {
		_ = os.Remove(data.Picture)
	}

	_, err = s.repo.DeleteVehicle(id)

	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	response := map[string]string{"message": "Vehicle deleted successfully"}

	return libs.GetResponse(response, 200, false)

}
