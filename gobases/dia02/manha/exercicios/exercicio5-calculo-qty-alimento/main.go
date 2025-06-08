package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func qtyDog(qty float64) float64 {
	return qty * 10.0
}

func qtyCat(qty float64) float64 {
	return qty * 5.0
}

func qtyHamster(qty float64) float64 {
	return qty * 0.25
}

func qtyTarantula(qty float64) float64 {
	return qty * 0.15
}

func Animal(animal string) (func(float64) float64, error) {
	switch animal {
	case dog:
		return qtyDog, nil
	case cat:
		return qtyCat, nil
	case hamster:
		return qtyHamster, nil
	case tarantula:
		return qtyTarantula, nil
	default:
		return nil, errors.New("operacao invalida")
	}
}

func main() {
	qtyDog, msg := Animal(dog)
	if msg != nil {
		fmt.Println(msg)
	}
	qtyCat, msg := Animal(cat)
	if msg != nil {
		fmt.Println(msg)
	}

	qtyHamster, msg := Animal(hamster)
	if msg != nil {
		fmt.Println(msg)
	}

	qtyTarantula, msg := Animal(tarantula)
	if msg != nil {
		fmt.Println(msg)
	}

	amountDog := qtyDog(5)
	amountCat := qtyCat(2)
	amountHamster := qtyHamster(8)
	amountTarantula := qtyTarantula(4)

	fmt.Printf("Comida total necessária para os cachorros: %.2f kg\n", amountDog)
	fmt.Printf("Comida total necessária para os gatos: %.2f kg\n", amountCat)
	fmt.Printf("Comida total necessária para os hamsters: %.2f kg\n", amountHamster)
	fmt.Printf("Comida total necessária para as tarantulas: %.2f kg\n", amountTarantula)
}
