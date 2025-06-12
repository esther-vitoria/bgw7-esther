package main

import (
	"fmt"
	"os"
)

func main() {
	// Cada função main começa pelo deferral da mensagem final.
	defer fmt.Println("execução concluída")

	_, err := os.Open("customer.txt")
	// defer file.Close()

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
}
