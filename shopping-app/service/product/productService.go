package product

type service struct {
	repo Repository
}

type Service interface {
	Create(p Product) (Product, error)
	GetAll() ([]Product, error)
	GetByID(id string) (Product, error)
	Update(id string, p Product) (Product, error)
	Delete(id string) error
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Create(p Product) (Product, error) {
	return s.repo.Create(p)
}

func (s *service) GetAll() ([]Product, error) {
	return s.repo.GetAll()
}

func (s *service) GetByID(id string) (Product, error) {
	return s.repo.GetByID(id)
}

func (s *service) Update(id string, p Product) (Product, error) {
	return s.repo.Update(id, p)
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}
