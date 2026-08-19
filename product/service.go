package product

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Get(code string) (Product, error) {
	return s.repo.Find(code)
}
