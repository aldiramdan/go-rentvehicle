package categories

import (
	"github.com/aldiramdan/go-backend/databases/orm/models"
	"github.com/aldiramdan/go-backend/interfaces"
	"github.com/aldiramdan/go-backend/libs"
	"gorm.io/gorm"
)

type cateogry_service struct {
	repo interfaces.CategoryRepo
}

func NewSrvc(repo interfaces.CategoryRepo) *cateogry_service {

	return &cateogry_service{repo}

}

func (s *cateogry_service) GetAllCategories() *libs.Response {

	result, err := s.repo.GetAllCategories()

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)
}

func (s *cateogry_service) GetCategoryById(id uint64) *libs.Response {

	result, err := s.repo.GetCategoryById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	return libs.GetResponse(result, 200, false)

}

func (s *cateogry_service) SearchCategories(query string) *libs.Response {

	result, err := s.repo.SearchCategories(query)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *cateogry_service) AddCategory(data *models.Category) *libs.Response {

	result, err := s.repo.AddCategory(data)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *cateogry_service) UpdateCategory(data *models.Category, id uint64) *libs.Response {

	_, err := s.repo.GetCategoryById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	result, err := s.repo.UpdateCategory(data, id)

	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	return libs.GetResponse(result, 200, false)

}

func (s *cateogry_service) DeleteCategory(id uint64) *libs.Response {

	_, err := s.repo.GetCategoryById(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	_, err = s.repo.DeleteCategory(id)

	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	response := map[string]string{"message": "Vehicle deleted successfully"}

	return libs.GetResponse(response, 200, false)

}
