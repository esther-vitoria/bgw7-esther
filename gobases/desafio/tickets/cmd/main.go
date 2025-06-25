package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	//Exemplo de uso

	var destinationInput string
	fmt.Printf("Insira um destino: ")
	fmt.Scanln(&destinationInput) //chamando um scan para inserir o destino

	total, err := tickets.GetTotalTickets(destinationInput) // chamando a func GetTotalTickets para pegar o total de tickets baseado no input do scan
	if err == nil {
		fmt.Printf("Total de tickets para o destino %s: %v", destinationInput, total)
	}

	allTickets, _ := tickets.LoadTickets() // chamando a func que carrega os dados do csv

	percent, err := tickets.PercentageDestination(destinationInput, len(allTickets)) //chamando a func AverageDestination que recebe o destino inserido e o total de tikets devolvido pela func LoadTickets

	if err == nil { //verificando se houve erro na chamada da func AverageDestination
		fmt.Printf("\nPorcentagem para o destino %s: %.2f%%", destinationInput, percent)
	}

	var periodInput string
	fmt.Printf("\nInsira um periodo (madrugada, manhã, tarde, noite): ")
	fmt.Scanln(&periodInput) //chamando um scan para inserir o periodo

	count, err := tickets.GetCountByPeriod(periodInput) //chamando a func GetCountByPeriod que recebe o periodo e retorna a quantidade dentro do periodo
	if err == nil {
		fmt.Printf("Quantidade no periodo %s: %v", periodInput, count)
	}

}
