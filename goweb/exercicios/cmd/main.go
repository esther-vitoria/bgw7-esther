package main

import (
	"net/http"

	"google.com/bgw7-esther/first-server/internal/application"
)

func main() {
	routes := application.NewRouter()

	if err := http.ListenAndServe(":8080", routes.MapRoutes()); err != nil {
		panic(err)
	}
}
