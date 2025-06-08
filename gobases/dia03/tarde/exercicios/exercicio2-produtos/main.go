package main

import "fmt"

type Product interface {
	Price() float64
}

type produtoPequeno struct {
	price float64
}

type produtoMedio struct {
	price float64
}

type produtoGrande struct {
	price float64
}

func (p produtoPequeno) Price() float64 {
	return p.price
}

func (p produtoMedio) Price() float64 {
	basePrice := p.price
	productTax := basePrice * 0.03
	storageTax := productTax * 0.03
	total := basePrice + productTax + storageTax
	return total
}

func (p produtoGrande) Price() float64 {
	basePrice := p.price
	productTax := basePrice * 0.06
	shippingCost := 2500.0
	total := basePrice + productTax + shippingCost
	return total
}

func NovoProduto(t string, price float64) Product {
	switch t {
	case "pequeno":
		return produtoPequeno{price}
	case "medio":
		return produtoMedio{price}
	case "grande":
		return produtoGrande{price}
	default:
		return nil
	}
}

func main() {
	p1 := NovoProduto("pequeno", 1000)
	p2 := NovoProduto("medio", 1000)
	p3 := NovoProduto("grande", 1000)

	fmt.Printf("Produto pequeno:  R$ %.2f\n", p1.Price())
	fmt.Printf("Produto médio: R$ %.2f\n", p2.Price())
	fmt.Printf("Produto grande: R$ %.2f\n", p3.Price())
}
