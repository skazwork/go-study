package product

import (
	"errors"
	"testing"
)

type fakeRepository struct {
	product Product
	err     error
}

func (r fakeRepository) Find(code string) (Product, error) {
	return r.product, r.err
}

func TestServiceGetSuccess(t *testing.T) {
	repo := fakeRepository{
		product: NewProduct("mouse", 5000, 2),
		err:     nil,
	}

	service := NewService(repo)
	p, err := service.Get("mouse")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if p.Name != "mouse" {
		t.Errorf("expected Mouse, got %s", p.Name)
	}
}

func TestServiceGetNotFound(t *testing.T) {
	repo := fakeRepository{
		product: Product{},
		err:     ErrNotFound,
	}

	service := NewService(repo)
	_, err := service.Get("Mouse")
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
	if err == nil {
		t.Errorf("expected ErrNotFound, got no errors")
	}
}

func TestServiceGetOutOfStock(t *testing.T) {
	repo := fakeRepository{
		product: NewProduct("iPhone", 100000, 0),
		err:     nil,
	}

	service := NewService(repo)
	_, err := service.Get("iPhone")

	if !errors.Is(err, ErrOutOfStock) {
		t.Errorf("expected out of stock error, got %v", repo.err)
	}
	if err == nil {
		t.Errorf("expected ErrOutOfStock, got no errors")
	}
}
