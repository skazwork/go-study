package main

import (
	"errors"
	"fmt"
	"go-study/product"
)

func main() {
	memRep := product.NewMemoryRepository()
	p, err := memRep.Find("mouse")

	if errors.Is(err, product.ErrNotFound) {
		fmt.Println("404")
		return
	}

	fmt.Println(p.Price())
}
