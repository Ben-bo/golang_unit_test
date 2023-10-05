package repository

import (
	"github.com/Ben-bo/golang_unit_test/entity"
	"github.com/stretchr/testify/mock"
)

type MockCategoryRepository struct {
	Mock mock.Mock
}

func (mock *MockCategoryRepository) findById(Id string) *entity.Category {
	arguments := mock.Mock.Called(Id)

	if arguments.Get(0) == nil {
		return nil
	}

	category := arguments.Get(0).(entity.Category)
	return &category
}
