package product

import "testing"

func TestProductTotal(t *testing.T) {
	p := NewProduct("Mouse", 5000, 3)
	if p.Total() != 15000 {
		t.Errorf("expected 15000, got %f", p.Total())
	}
}
