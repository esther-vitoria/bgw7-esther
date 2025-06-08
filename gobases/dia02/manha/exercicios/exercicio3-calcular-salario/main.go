package main

import "fmt"

const (
	categoriaC = "C"
	categoriaB = "B"
	categoriaA = "A"
)

func main() {

	resultado := calculoSalario(9600, categoriaB)

	fmt.Printf("O Salário tem o valor de:  %.2f", resultado)

}

func calculoSalario(minutosTrabalhados float64, categoria string) (salarioFinal float64) {

	switch categoria {
	case categoriaC:
		salarioFinal = opCategoriaC(minutosTrabalhados)
	case categoriaB:
		salarioFinal = opCategoriaB(minutosTrabalhados)
	case categoriaA:
		salarioFinal = opCategoriaA(minutosTrabalhados)

	}
	return salarioFinal
}

func opCategoriaC(minutosTrabalhados float64) (salarioFinal float64) {
	salarioHora := 1000.0
	horasTrabalhadas := minutosTrabalhados / 60
	salarioFinal = salarioHora * horasTrabalhadas
	return salarioFinal
}

func opCategoriaB(minutosTrabalhados float64) (salarioFinal float64) {
	salarioHora := 1500.0
	horasTrabalhadas := minutosTrabalhados / 60
	salarioMensal := salarioHora * horasTrabalhadas
	percentual20 := (salarioMensal * 20.0) / 100.0
	salarioFinal = salarioMensal + percentual20
	return salarioFinal
}

func opCategoriaA(minutosTrabalhados float64) (salarioFinal float64) {
	salarioHora := 3000.0
	horasTrabalhadas := minutosTrabalhados / 60
	salarioMensal := salarioHora * horasTrabalhadas
	percentual50 := (salarioMensal * 50.0) / 100.0
	salarioFinal = salarioMensal + percentual50
	return salarioFinal
}
