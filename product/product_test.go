package product

import "testing"

func TestProductTotal(t *testing.T) {
	p := NewProduct("Mouse", 5000, 3)
	total := p.Total()

	if total != 15000 {
		t.Errorf("expected 15000, got %f", total)
	}
}
