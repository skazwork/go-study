package product

import "errors"

type MemoryRepository struct {
	products map[string]Product
}

func NewMemoryRepository() *MemoryRepository {
	products := map[string]Product{
		"mouse":    NewProduct("Mouse", 4999.9, 2),
		"keyboard": NewProduct("Keyboa
		rd", 1000.0, 1),
	}
	return &MemoryRepository{
		products: products,
	}
}

func (r *MemoryRepository) Find(code string) (Product, error) {
	product, ok := r.products[code]
	if !ok {
		return Product{}, ErrNotFound
	}

	return product, nil
}
