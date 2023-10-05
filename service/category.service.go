package service

import (
	"errors"

	"github.com/Ben-bo/golang_unit_test/entity"
	"github.com/Ben-bo/golang_unit_test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service *CategoryService) GetCategory(Id string) (*entity.Category, error) {
	category := service.Repository.FindById(Id)

	if category == nil {
		return category, errors.New("category not found")
	}

	return category, nil

}
