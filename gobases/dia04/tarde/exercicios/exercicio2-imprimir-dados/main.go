package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("execução concluída")
	}()

	//Tenta abrir o arquivo
	file, err := os.Open("/Users/enbalbino/Documents/GitHub/bgw7-esther/gobases/dia04/tarde/exercicios/exercicio2-imprimir-dados/customer.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	defer file.Close()

	// Lê o conteúdo do arquivo e imprime
	content, err := os.ReadFile("/Users/enbalbino/Documents/GitHub/bgw7-esther/gobases/dia04/tarde/exercicios/exercicio2-imprimir-dados/customer.txt")
	if err != nil || string(content) == "" {
		panic("Error on trying read the data")
	}

	fmt.Println("File data:")
	fmt.Println(string(content))
}
