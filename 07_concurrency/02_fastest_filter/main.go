package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func readFile(fileName string) []Product {
	var err error
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var products []Product
	err = json.Unmarshal(content, &products)
	if err != nil {
		panic(err)
	}
	return products
}

func findCheapProducts(products []Product) []Product {
	var foundProducts []Product
	for _, product := range products {
		if product.Price < 50 {
			foundProducts = append(foundProducts, product)
		}
	}
	return foundProducts
}

func main() {
	products0 := readFile("products_0.json")
	products1 := readFile("products_1.json")
	products2 := readFile("products_2.json")

	// Sequentially
	start := time.Now()
	cheapProducts := findCheapProducts(products0)
	cheapProducts = append(cheapProducts, findCheapProducts(products1)...)
	cheapProducts = append(cheapProducts, findCheapProducts(products2)...)
	fmt.Println(len(cheapProducts))
	fmt.Println("Took", time.Now().Sub(start))

	// Concurrent
	start = time.Now()
	// TODO: Implement concurrent use of findCheapProducts on products0, products1 and product3
	fmt.Println(len(cheapProducts))
	fmt.Println("Took", time.Now().Sub(start))
}
