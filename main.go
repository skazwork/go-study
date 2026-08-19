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
	repo := product.NewMemoryRepository(products)
	service := product.NewService(repo)

	findAndPrint(*service, "mouse")
}

func findAndPrint(serv product.Service, code string) {
	p, err := serv.Get("mouse")
	if errors.Is(err, product.ErrNotFound) {
		fmt.Println("404")
		return
	} else if err != nil {
		fmt.Println("fatal error")
		return
	}

	fmt.Println(p.Price())
}
