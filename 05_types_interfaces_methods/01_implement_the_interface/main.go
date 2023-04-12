package main

import (
	"fmt"
	"reflect"
)

type Product struct {
	ID    string
	Name  string
	Price int64
}

type ProductStorage interface {
	Add(product Product)
	Remove(productId string)
	List() []Product
}

// TODO: Make RealProductStorage implement the interface ProductStorage
type RealProductStorage struct{}

func NewRealProductStorage() *RealProductStorage {
	panic("Implement me")
}

func main() {
	realProductStorage := NewRealProductStorage()
	realProductStorage.Add(Product{
		ID:    "mentosMint",
		Name:  "Mentos Mint",
		Price: 199,
	})
	realProductStorage.Add(Product{
		ID:    "mentosFruit",
		Name:  "Mentos Fruit",
		Price: 149,
	})

	if !reflect.DeepEqual(realProductStorage.List(), []Product{
		{
			ID:    "mentosMint",
			Name:  "Mentos Mint",
			Price: 199,
		},
		{
			ID:    "mentosFruit",
			Name:  "Mentos Fruit",
			Price: 149,
		},
	}) {
		panic("Wrong")
	} else {
		fmt.Println("Correct ✅")
	}

	realProductStorage.Remove("mentosMint")

	if !reflect.DeepEqual(realProductStorage.List(), []Product{
		{
			ID:    "mentosFruit",
			Name:  "Mentos Fruit",
			Price: 149,
		},
	}) {
		panic("Wrong")
	} else {
		fmt.Println("Correct ✅")
	}
}
