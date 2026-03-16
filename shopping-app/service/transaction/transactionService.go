package transaction

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type service struct {
	repo Repository
}

type Service interface {
	Create(t Transaction) (Transaction, error)
	GetAll() ([]Transaction, error)
	GetByID(id string) (Transaction, error)
	Update(id string, t Transaction) (Transaction, error)
	Delete(id string) error
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Create(t Transaction) (Transaction, error) {

	payment := map[string]interface{}{
		"transaction_id": t.ID.Hex(),
		"amount":         t.Total,
	}

	body, _ := json.Marshal(payment)

	resp, err := http.Post(
		"http://payment-app:8081/payments",
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return t, err
	}

	if resp.StatusCode != http.StatusCreated {
		return t, errors.New("payment failed")
	}

	t.Status = "paid"

	return s.repo.Create(t)
}

func (s *service) GetAll() ([]Transaction, error) {
	return s.repo.GetAll()
}

func (s *service) GetByID(id string) (Transaction, error) {
	return s.repo.GetByID(id)
}

func (s *service) Update(id string, t Transaction) (Transaction, error) {
	return s.repo.Update(id, t)
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}
