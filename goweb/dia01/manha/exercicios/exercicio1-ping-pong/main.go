package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter() //Instaciando um router para controlar os endpoints

	//Criando uma rota GET
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("pong")) //devolvendo a palavra "pong"

	})

	//incializa o servidor HTTP escutando a porta 8080
	//o router vai lidar com todas as requisições
	http.ListenAndServe(":8080", r)

}
