package product

type MemoryRepository struct {
	products map[string]Product
}

func NewMemoryRepository(products map[string]Product) *MemoryRepository {
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
