package transaction

type Repository interface {
	Create(t Transaction) (Transaction, error)
	GetAll() ([]Transaction, error)
	GetByID(id string) (Transaction, error)
	Update(id string, t Transaction) (Transaction, error)
	Delete(id string) error
}
