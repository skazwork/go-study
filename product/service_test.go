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

func TestServiceGet(t *testing.T) {
	tests := []struct {
		name        string
		repoProduct Product
		repoErr     error
		code        string
		wantName    string
		wantErr     error
	}{
		{
			name:        "success",
			repoProduct: NewProduct("Mouse", 5000, 2),
			repoErr:     nil,
			code:        "mouse",
			wantName:    "Mouse",
			wantErr:     nil,
		},
		{
			name:        "not found",
			repoProduct: Product{},
			repoErr:     ErrNotFound,
			code:        "iphone",
			wantErr:     ErrNotFound,
		},
		{
			name:        "out of stock",
			repoProduct: NewProduct("iPhone", 100000, 0),
			repoErr:     nil,
			code:        "iphone",
			wantErr:     ErrOutOfStock,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := fakeRepository{
				product: tt.repoProduct,
				err:     tt.repoErr,
			}

			service := NewService(repo)

			p, err := service.Get(tt.code)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("expected error %v, got %v", tt.wantErr, err)
				}
			} else {
				if p.Name != tt.wantName {
					t.Errorf("expected name %s, got %s", tt.wantName, p.Name)
				}
			}
		})
	}
}
