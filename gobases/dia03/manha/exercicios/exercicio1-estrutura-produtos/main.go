package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var products = []Product{
	{ID: 1, Name: "Camiseta", Price: 49.90, Description: "Camiseta de algodão", Category: "Vestuário"},
	{ID: 2, Name: "Fone de Ouvido", Price: 199.90, Description: "Fone bluetooth", Category: "Eletrônicos"},
}

// Método Save: adiciona o produto ao slice global products
func (p Product) Save() {
	products = append(products, p)
}

// Método GetAll: imprime todos os produtos salvos
func (p Product) GetAll() {
	fmt.Println("Produtos salvos:")
	for _, prod := range products {
		fmt.Printf("ID: %d | Nome: %s | Preço: %.2f | Descrição: %s | Categoria: %s\n",
			prod.ID, prod.Name, prod.Price, prod.Description, prod.Category)
	}
}

// Função getById: busca produto pelo id e retorna ele e um erro (caso não exista)
func getById(id int) (Product, error) {
	for _, prod := range products {
		if prod.ID == id {
			return prod, nil
		}
	}
	return Product{}, errors.New("produto não encontrado")
}

func main() {

	novoProduto := Product{ID: 3, Name: "Notebook", Price: 3500.00, Description: "Notebook gamer", Category: "Informática"}
	novoProduto.Save()

	novoProduto.GetAll()

	prod, err := getById(3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nProduto encontrado por ID %v :\nNome: %s | Preço: %.2f\n", prod.ID, prod.Name, prod.Price)
	}
}
