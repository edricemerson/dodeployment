package product

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(p Product) (Product, error) {
	args := m.Called(p)
	return args.Get(0).(Product), args.Error(1)
}

func (m *MockRepository) GetAll() ([]Product, error) {
	args := m.Called()
	return args.Get(0).([]Product), args.Error(1)
}

func (m *MockRepository) GetByID(id string) (Product, error) {
	args := m.Called(id)
	return args.Get(0).(Product), args.Error(1)
}

func (m *MockRepository) Update(id string, p Product) (Product, error) {
	args := m.Called(id, p)
	return args.Get(0).(Product), args.Error(1)
}

func (m *MockRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
