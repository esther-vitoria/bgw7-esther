package main

import (
	"errors"
	"fmt"
)

type Customer struct {
	ID      string
	Name    string
	Tel     string
	Address string
}

// Banco em memória
var Customers = []Customer{
	{ID: "001", Name: "Esther Vitoria", Tel: "51999999999", Address: "Av. Mercado Livre, 123"},
}

func main() {
	defer func() {
		fmt.Println("End of execution")
		if r := recover(); r != nil {
			fmt.Println("Several errors were detected at runtime", r)
		}
	}()

	// Tentando adicionar um Customer já existente
	newCustomer := Customer{ID: "001", Name: "João Santana", Tel: "51999999999", Address: "Av. Principal, 123"}
	registerCustomer(newCustomer)

	// Tentando adicionar um cliente novo com campo vazio (deve dar erro)
	newCustomer2 := Customer{ID: "003", Name: "", Tel: "51990000000", Address: "Avenida Sete"}
	registerCustomer(newCustomer2)
}

func registerCustomer(c Customer) {
	// 1. Verifica se já existe
	for _, Customer := range Customers {
		if Customer.ID == c.ID {
			// Imprime mensagem de erro com panic, mas permite continuar
			fmt.Println("Error: Customer already exists")
			panic("Erro: Customer already exists")
		}
	}

	// 2. Validação dos dados não serem zero
	if ok, err := checkCustomer(c); !ok {
		panic(err)
	}

	// 3. Se tudo ok, adiciona
	Customers = append(Customers, c)
	fmt.Println("Customer sucessfuly registered:", c)
}

// Retorna se está tudo ok, se não, retorna erro personalizado
func checkCustomer(c Customer) (bool, error) {
	if c.ID == "" || c.Name == "" || c.Tel == "" || c.Address == "" {
		return false, errors.New("All fields need to be filled in!")
	}
	return true, nil
}
