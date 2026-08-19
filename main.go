package main

import (
	"errors"
	"fmt"
	"go-study/product"
)

func main() {
	products := map[string]product.Product{
		"mouse":    product.NewProduct("Mouse", 4999.9, 2),
		"keyboard": product.NewProduct("Keyboard", 1000.0, 1),
		"iphone":   product.NewProduct("iPhone", 100000, 0),
	}
	repo := product.NewMemoryRepository(products)
	service := product.NewService(repo)

	findAndPrint(service, "iphone")
}

func findAndPrint(serv product.ProductGetter, code string) {
	p, err := serv.Get(code)
	if err != nil {
		processErrors(err)
		return
	}

	fmt.Println(p.Price())
}

func processErrors(err error) {
	switch {
	case errors.Is(err, product.ErrNotFound):
		{
			fmt.Println("404")
		}
	case errors.Is(err, product.ErrOutOfStock):
		{
			fmt.Println("409: product out of stock")
		}
	default:
		{
			fmt.Println("fatal error")
		}
	}
}
