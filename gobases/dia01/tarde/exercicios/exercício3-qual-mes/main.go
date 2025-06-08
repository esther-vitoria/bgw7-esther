package main

import "fmt"

// Para esse exercício foi feito um array que contém todos os meses do ano, a solução recebe um valor com o numero do mês, faz o decremento desse valor para que
// seja acessado o índice correto dentro do array, também é feito uma checkagem para verificar se o número de entrada é um mês válido
func main() {
	numeroMes := 7
	meses := []string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}
	numeroMes--

	if numeroMes >= 0 && numeroMes < 12 {
		fmt.Println(meses[numeroMes])
	} else {
		fmt.Print("Número invalido!")
	}

}
