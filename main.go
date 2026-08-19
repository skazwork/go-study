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
	}

	p, err := product.Find(products, "mouse")

	if errors.Is(err, product.ErrNotFound) {
		fmt.Println("404")
		return
	}

	fmt.Println(p.Price())
}
