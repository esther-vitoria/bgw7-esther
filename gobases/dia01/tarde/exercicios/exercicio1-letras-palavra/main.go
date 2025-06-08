package main

import "fmt"

func main() {
	palavra := "Borboleta" // Utilizando a declaração de variável resumida e atribuíndo o valor Borboleta

	fmt.Println("Número de letras: ", len(palavra)) // Aqui estou imprimindo o quantas letras a palavra tem utilizando a função len(variavel)

	for _, letra := range palavra { //Nesse for estou usando for seguido do underscore para percorrer todos os caractéres da string, ou seja, ele irá percorrer todos os caractéres até o fim
		fmt.Printf("%c\n", letra) // Aqui estou utilizando o Printf %c que imprime o caractere como texto (e não o código Unicode)
	}
}
