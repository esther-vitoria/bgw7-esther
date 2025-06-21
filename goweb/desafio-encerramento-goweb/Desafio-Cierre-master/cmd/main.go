package main

import (
	"app/desafio-goweb/internal/application"
	"fmt"
)

func main() {

	cfg := &application.ConfigAppDefault{
		ServerAddr: ":8080",
		DbFilePath: "docs/db/tickets.csv",
	}

	app := application.NewServerChi(cfg)

	err := app.SetUp()
	if err != nil {
		fmt.Println("Iniciando....")
		fmt.Println(err)
		return
	}
}
