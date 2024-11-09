package product

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllProducts() []Product {
	return s.repo.GetAll()
}

func (s *Service) GetProductByID(id int) (*Product, error) {
	return s.repo.GetByID(id)
}

func (s *Service) CreateProduct(product Product) Product {
	return s.repo.Create(product)
}

func (s *Service) UpdateProduct(id int, product Product) (*Product, error) {
	return s.repo.Update(id, product)
}

func (s *Service) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}
