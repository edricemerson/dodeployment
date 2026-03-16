package payment

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(p Payment) (Payment, error) {

	args := m.Called(p)

	return args.Get(0).(Payment), args.Error(1)
}

func (m *MockRepository) ReadByTransactionID(txID string) (Payment, error) {

	args := m.Called(txID)

	return args.Get(0).(Payment), args.Error(1)
}
