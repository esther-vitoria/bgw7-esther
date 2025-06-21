package storage

import (
	"encoding/json"
	"os"
)

// struct para representar os dados arquivo JSON
type ProductJSON struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// Struct que pega os dados do body da requisição
type RequestBodyProduct struct {
	Name        *string  `json:"name"`
	Quantity    *int     `json:"quantity"`
	CodeValue   *string  `json:"code_value"`
	IsPublished *bool    `json:"is_published"`
	Expiration  *string  `json:"expiration"`
	Price       *float64 `json:"price"`
}

// func para carregar os produtos do arquivo JSON
func LoadProducts(path string) []*ProductJSON {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var products []*ProductJSON
	_ = json.Unmarshal(data, &products)
	return products
}
