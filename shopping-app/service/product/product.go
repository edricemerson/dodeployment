package product

type Repository interface {
	Create(p Product) (Product, error)
	GetAll() ([]Product, error)
	GetByID(id string) (Product, error)
	Update(id string, p Product) (Product, error)
	Delete(id string) error
}
