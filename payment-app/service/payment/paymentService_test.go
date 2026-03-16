package payment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePaymentSuccess(t *testing.T) {

	mockRepo := new(MockRepository)

	service := NewService(mockRepo)

	input := Payment{
		TransactionID: "TX1001",
		Amount:        50000,
	}

	expected := Payment{
		TransactionID: "TX1001",
		Amount:        50000,
		Status:        "success",
	}

	mockRepo.On("Create", expected).Return(expected, nil)

	result, err := service.Create(input)

	assert.NoError(t, err)
	assert.Equal(t, "success", result.Status)

	mockRepo.AssertExpectations(t)
}

func TestCreatePaymentInvalidAmount(t *testing.T) {

	mockRepo := new(MockRepository)

	service := NewService(mockRepo)

	input := Payment{
		TransactionID: "TX1001",
		Amount:        -10,
	}

	_, err := service.Create(input)

	assert.Error(t, err)
	assert.Equal(t, "invalid payment amount", err.Error())
}
