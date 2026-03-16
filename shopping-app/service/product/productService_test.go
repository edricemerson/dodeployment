package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {

	mockRepo := new(MockRepository)

	input := Product{
		Name:  "Laptop",
		Price: 1200,
		Stock: 5,
	}

	mockRepo.On("Create", input).Return(input, nil)

	service := NewService(mockRepo)

	result, err := service.Create(input)

	assert.NoError(t, err)
	assert.Equal(t, input, result)

	mockRepo.AssertExpectations(t)
}

func TestGetAllProducts(t *testing.T) {

	mockRepo := new(MockRepository)

	expected := []Product{
		{Name: "Laptop", Price: 1200, Stock: 5},
		{Name: "Mouse", Price: 20, Stock: 50},
	}

	mockRepo.On("GetAll").Return(expected, nil)

	service := NewService(mockRepo)

	result, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	mockRepo.AssertExpectations(t)
}

func TestGetProductByID(t *testing.T) {

	mockRepo := new(MockRepository)

	expected := Product{
		Name:  "Keyboard",
		Price: 100,
		Stock: 10,
	}

	mockRepo.On("GetByID", "123").Return(expected, nil)

	service := NewService(mockRepo)

	result, err := service.GetByID("123")

	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	mockRepo.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {

	mockRepo := new(MockRepository)

	input := Product{
		Name:  "Monitor",
		Price: 300,
		Stock: 7,
	}

	mockRepo.On("Update", "123", input).Return(input, nil)

	service := NewService(mockRepo)

	result, err := service.Update("123", input)

	assert.NoError(t, err)
	assert.Equal(t, input, result)

	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {

	mockRepo := new(MockRepository)

	mockRepo.On("Delete", "123").Return(nil)

	service := NewService(mockRepo)

	err := service.Delete("123")

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
