package product

import "errors"

type Repository interface {
	Find(code string) (Product, error)
}

var ErrNotFound = errors.New("product not found")
