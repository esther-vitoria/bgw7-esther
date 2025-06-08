package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
)

// struct to represent the json fields
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// Check the JSON File
func ShowProducts(filepath string) ([]Product, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var products []Product
	err = json.Unmarshal(data, &products)
	return products, err
}

// Respond all products listed on JSON file
func GetProductHandler(products []Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

// Respond the product based on id
func GetProductByIdHandler(products []Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		for _, p := range products {
			if p.ID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(p)
				return
			}
		}
		http.NotFound(w, r)

	}
}

func SearchProductsHandler(products []Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		priceGtStr := r.URL.Query().Get("priceGt")
		if priceGtStr == "" {
			http.Error(w, "priceGt param is required", http.StatusBadRequest)
			return
		}
		priceGt, err := strconv.ParseFloat(priceGtStr, 64)
		if err != nil {
			http.Error(w, "priceGt must be a number", http.StatusBadRequest)
			return
		}
		var result []Product
		for _, p := range products {
			if p.Price > priceGt {
				result = append(result, p)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
