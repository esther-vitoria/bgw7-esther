package main

import (
	"encoding/json"
	"log"
	"os"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Qty         int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func ShowProducts(filename string) ([]Product, error) {
	datas, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var products []Product
	err = json.Unmarshal(datas, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func main() {
	products, err := ShowProducts("products.json")
	if err != nil {
		log.Fatal("Error loading products:", err)
	}
	log.Println("Loaded products:", products)

}
