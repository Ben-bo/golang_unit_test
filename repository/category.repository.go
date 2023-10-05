package repository

import "github.com/Ben-bo/golang_unit_test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
