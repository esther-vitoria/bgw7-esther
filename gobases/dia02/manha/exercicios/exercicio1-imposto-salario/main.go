package main

import "fmt"

func main() {
	var salario float64
	fmt.Printf("Insira o valor do salário: ")
	fmt.Scanln(&salario)

	impostoSalario(salario)
}

func impostoSalario(salario float64) (deducao float64) {
	var deducaoPercentual float64

	if salario >= 50000.0 && salario < 150000.0 {
		deducaoPercentual = 0.17
	} else if salario >= 150000.0 {
		deducaoPercentual = 0.27
	} else {
		deducaoPercentual = 0.10
	}

	deducao = salario * deducaoPercentual
	salarioDeduzido := salario - deducao

	fmt.Printf("Valor incial do salário: %.2f \n A dedução foi de %.0f%% \n O valor deduzido foi: %.2f \n O salário atual é de: %.2f", salario, deducaoPercentual*100, deducao, salarioDeduzido)

	return
}
