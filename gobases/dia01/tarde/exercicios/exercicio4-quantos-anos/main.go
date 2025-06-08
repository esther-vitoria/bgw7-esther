package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Map inicial: ", employees)

	fmt.Println("Imprimindo a idade de Benjamin: ", employees["Benjamin"]) //Imprime a idade de Benjamin

	qtyFuncionario := 0
	for _, element := range employees { // for que verifica a idade dos funcionários e imprime quantos tem idade superior de 21 anos

		if element > 21 {
			qtyFuncionario++
		}
	}
	fmt.Println("Quantidade de funcionários acima de 21 anos: ", qtyFuncionario)

	employees["Federico"] = 25 // Adicionando Federico de 25 anos a lista
	fmt.Println("Adicionando Federico a lista ", employees)

	delete(employees, "Pedro") // Deletando Pedro da lista
	fmt.Println("Deletando Pedro da lista: ", employees)
}
