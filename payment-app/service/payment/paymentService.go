package payment

import "errors"

type service struct {
	repo Repository
}

type Service interface {
	Create(p Payment) (Payment, error)
	GetByTransactionID(txID string) (p Payment, err error)
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) Create(p Payment) (Payment, error) {

	if p.Amount <= 0 {
		return p, errors.New("invalid payment amount")
	}

	p.Status = "success"

	return s.repo.Create(p)
}

func (s *service) GetByTransactionID(txID string) (p Payment, err error) {
	return s.repo.ReadByTransactionID(txID)
}
