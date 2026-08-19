package product

type ProductGetter interface {
	Get(code string) (Product, error)
}
