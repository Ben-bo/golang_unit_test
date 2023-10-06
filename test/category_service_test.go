package test

import (
	"testing"

	"github.com/Ben-bo/golang_unit_test/entity"
	"github.com/Ben-bo/golang_unit_test/repository"
	"github.com/Ben-bo/golang_unit_test/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.MockCategoryRepository{Mock: mock.Mock{}}
var categoryService = service.CategoryService{Repository: categoryRepository}

func TestCategoryServiceGetNotFound(t *testing.T) {
	//programm mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.GetCategory("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}
func TestCategoryServiceGetSuccess(t *testing.T) {
	dataCategory := entity.Category{
		Id:   "2",
		Name: "test",
	}

	//programm mock
	categoryRepository.Mock.On("FindById", "2").Return(dataCategory)

	category, err := categoryService.GetCategory("2")

	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.ObjectsAreEqual(dataCategory, category)
}
