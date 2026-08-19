package product

import "errors"

var ErrNotFound = errors.New("product not found")

func Find(products map[string]Product, code string) (Product, error) {
	product, ok := products[code]
	if !ok {
		return Product{}, ErrNotFound
	}

	return product, nil
}
