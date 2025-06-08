package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {

	r := chi.NewRouter()

	r.Post("/grettings", func(w http.ResponseWriter, r *http.Request) {
		var p Person

		//decodifica o json recebido
		err := json.NewDecoder(r.Body).Decode(&p)

		if err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		msg := fmt.Sprintf("Hello %s %s, welcome to meli!! :)", p.FirstName, p.LastName)

		//retornando o texto
		w.Write([]byte(msg))

	})
	http.ListenAndServe(":8080", r)

}
