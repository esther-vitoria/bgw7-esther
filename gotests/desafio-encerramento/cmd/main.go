package main

import (
	"app/internal/application"
	"fmt"
)

func main() {
	cfg := &application.ConfigApplicationDefault{
		Addr:            "127.0.0.1:8080",
		ProductJSONPath: "database/products.json",
	}

	app := application.NewApplicationDefault(cfg)
	defer app.TearDown()
	if err := app.SetUp(cfg); err != nil {
		fmt.Println("Erro no setup:", err)
		return
	}

	fmt.Printf("✅ Servidor rodando em %s\n", cfg.Addr)

	if err := app.Run(); err != nil {
		fmt.Println("Erro no run:", err)
		return
	}
}
