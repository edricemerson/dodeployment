package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTransactions(t *testing.T) {

	mockRepo := new(MockRepository)

	expected := []Transaction{
		{Quantity: 2, Total: 200, Status: "paid"},
	}

	mockRepo.On("GetAll").Return(expected, nil)

	service := NewService(mockRepo)

	result, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	mockRepo.AssertExpectations(t)
}

func TestGetTransactionByID(t *testing.T) {

	mockRepo := new(MockRepository)

	expected := Transaction{
		Quantity: 1,
		Total:    100,
		Status:   "paid",
	}

	mockRepo.On("GetByID", "123").Return(expected, nil)

	service := NewService(mockRepo)

	result, err := service.GetByID("123")

	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	mockRepo.AssertExpectations(t)
}

func TestDeleteTransaction(t *testing.T) {

	mockRepo := new(MockRepository)

	mockRepo.On("Delete", "123").Return(nil)

	service := NewService(mockRepo)

	err := service.Delete("123")

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
