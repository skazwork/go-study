package product

type Product struct {
	Name  string
	price float64
	Qty   int
}

func (p Product) Total() float64 {
	return p.price * float64(p.Qty)
}

func (p Product) Price() float64 {
	return p.price
}

func NewProduct(name string, price float64, qty int) Product {
	return Product{
		Name:  name,
		price: price,
		Qty:   qty,
	}
}
