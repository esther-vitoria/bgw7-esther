package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
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
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	})

	fmt.Println("API rodando em http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", r))

}
