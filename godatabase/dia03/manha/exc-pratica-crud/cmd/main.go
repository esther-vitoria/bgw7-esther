package main

import (
	"net/http"

	"github.com/bgw7/exc-pratica-crud/cmd/router"
	"github.com/bgw7/exc-pratica-crud/config"
	"github.com/bgw7/exc-pratica-crud/database"
)

func main() {
	config.LoadConfig()
	config.NewConfig(nil)

	err := database.Connect()
	if err != nil {
		panic(err)
	}

	defer database.GetDB().Close()

	r := router.NewRouter()

	if err := http.ListenAndServe(":8080", r.MapRoutes()); err != nil {
		panic(err)
	}
}
