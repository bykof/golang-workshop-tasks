package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Rating struct {
	Rate  float64 `json:"rate"`
	Count int64   `json:"count"`
}

type Product struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      Rating  `json:"rating"`
}

// TODO: Implement the interface for RealProductService and FakeProductService called "ProductService".

type RealProductService struct {
	baseUrl string
	client  *http.Client
}

func (ps *RealProductService) Products() ([]Product, error) {
	response, err := ps.client.Get(fmt.Sprintf("%s/products", ps.baseUrl))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	var products []Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}
	return products, err
}

func (ps *RealProductService) Product(id int64) (Product, error) {
	response, err := ps.client.Get(fmt.Sprintf("%s/products/%d", ps.baseUrl, id))
	if err != nil {
		return Product{}, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		return Product{}, err
	}
	return product, err
}

func NewRealProductService(baseUrl string) *RealProductService {
	return &RealProductService{
		baseUrl: baseUrl,
		client:  &http.Client{},
	}
}

type FakeProductService struct{}

// TODO: Implement other functions

func NewFakeProductService() *FakeProductService {
	return &FakeProductService{}
}

func main() {
	var productService ProductService
	productService = NewRealProductService("https://fakestoreapi.com")
	products, err := productService.Products()
	if err != nil {
		panic(err)
	}

	for _, product := range products[:3] {
		product, err := productService.Product(product.ID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n\n", product)
	}

	productService = NewFakeProductService()
	products, err = productService.Products()
	if err != nil {
		panic(err)
	}

	for _, product := range products[:3] {
		product, err := productService.Product(product.ID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n\n", product)
	}
}
