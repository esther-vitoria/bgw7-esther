package main

import "fmt"

func main() {

	//Coletando variáveis
	idade := 20
	empregado := true
	anosEmpregado := 5
	salario := 4.500

	if idade >= 22 && empregado == true && anosEmpregado > 1 && salario >= 100.000 { //Aqui estou verificando a condição de empréstimo sem juros e para isso o cliente precisa ganhar mais de 100.000
		fmt.Println("Seu empréstimo foi aprovado sem juros!")

	} else if idade >= 22 && empregado == true && anosEmpregado > 1 && salario < 100.000 { //Aqui estou verificando a condição de empréstimo com juros caso o cliente ganhe menos que 100.000
		fmt.Println("Seu empréstimo foi aprovado com juros!")

	} else { //Aqui é se caso o cliente não se encaixe em nenhuma da exceções ele não terá direito ao empréstimo
		fmt.Println("Seu empréstimo não foi aprovado!")
	}
}
