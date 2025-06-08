package main

import "fmt"

func main() {
	var firstName string      // Corrigido: nome válido, tipo depois do nome.
	var lastName string       // Já estava correto.
	var age int               // Corrigido: tipo depois do nome.
	var driver_license = true // Já estava correto.
	var personHeight int      // Corrigido: nome de variável único e em camelCase.
	childsNumber := 2         // Já estava correto; declaração curta, válida dentro de função.

	// Para utlizar as variáveis e não dar erro
	fmt.Println(firstName, lastName, age, driver_license, personHeight, childsNumber)
}

/*
var 1firstName string
Erro: Nomes de variáveis não podem começar com número.
Correto: var firstName string

var lastName string
Correto!

var int age
Erro: A ordem está incorreta; o tipo deve vir depois do nome da variável.
Correto: var age int

1lastName := 6
Erro: Nomes de variáveis não podem começar com um número; além disso, := só pode ser usado dentro de funções.
Correto: lastName := 6 (usando um nome válido e dentro da função)

var driver_license = true
Correto! Se atribuir true ou false o Go já entende que é uma variável do tipo boolean

var person height int
Erro: Não pode haver espaço no nome da variável; além disso, somente um nome para a variável antes do tipo.
Correto: var personHeight int

childsNumber := 2
Correto!
Nome válido, inicialização curta (desde que dentro de uma função).

*/
