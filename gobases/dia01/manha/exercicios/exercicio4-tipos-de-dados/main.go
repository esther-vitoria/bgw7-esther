package main

import "fmt"

func main() {
	var age int = 35
	boolean := false
	var salary float64 = 45857.90

	// Para utilizar as variáveis e não dar erro
	fmt.Print(age, boolean, salary)
}

/*
var lastName string = "Smith"
Correto! A variável segue o padrão (var nome tipo) e faz a atribuição de valores correspondentes ao tipo atribuido a variável

var age int = "35"
Incorreto!! Aqui a variável foi declarada corretamente mas o valor atribuido a ela está errado, o correto seria: var age int = 35

boolean := "false"
Incorreto!! Aqui a variável foi declarada corretamente mas o valor atribuido a ela está errado e ela precisa estar dentro de uma função para ser usada, o correto seria: boolean := false

var salary string = 45857.90
Incorreto!! Aqui a variável foi declarada corretamente mas o valor atribuido a ela está errado pois deveria ser do tipo float, o correto seria: var salary float64 = 45857.90

var firstName string = "Mary"
Correto! A variável segue o padrão (var nome tipo) e faz a atribuição de valores correspondentes ao tipo atribuido a variável
*/
