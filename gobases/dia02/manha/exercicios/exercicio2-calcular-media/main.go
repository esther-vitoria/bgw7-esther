package main

import "fmt"

func main() {
	calculoMedia(5, 10, 7, 8, 9, 6, 4)
}

func calculoMedia(notas ...int) (media int) {
	var soma int
	for _, nota := range notas {
		if nota < 0 {
			fmt.Printf("O número %d é negativo!", nota)
			panic("Verifique os números inserido!")
		} else {
			soma += nota
		}
	}

	media = soma / len(notas)
	fmt.Printf("A média das notas é %d!", media)
	return
}
