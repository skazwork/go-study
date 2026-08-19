package product

import "errors"

var ErrOutOfStock = errors.New("product out of stock")

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Get(code string) (Product, error) {
	p, err := s.repo.Find(code)
	if err != nil {
		return Product{}, err
	}
	if p.Qty <= 0 {
		return Product{}, ErrOutOfStock
	}

	return p, nil
}
