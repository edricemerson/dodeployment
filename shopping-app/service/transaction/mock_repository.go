package transaction

import (
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(t Transaction) (Transaction, error) {
	args := m.Called(t)
	return args.Get(0).(Transaction), args.Error(1)
}

func (m *MockRepository) GetAll() ([]Transaction, error) {
	args := m.Called()
	return args.Get(0).([]Transaction), args.Error(1)
}

func (m *MockRepository) GetByID(id string) (Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(Transaction), args.Error(1)
}

func (m *MockRepository) Update(id string, t Transaction) (Transaction, error) {
	args := m.Called(id, t)
	return args.Get(0).(Transaction), args.Error(1)
}

func (m *MockRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
