package middlewares

import (
	"net/http"
	"os"

	"google.com/bgw7-esther/first-server/internal/utils"
)

func TokenAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Access-Token")
		secret := os.Getenv("ACCESS_TOKEN")

		if token != secret || token == "" {
			utils.Respond(w, http.StatusUnauthorized, &utils.RespondBodyProduct{
				Message: "Invalid token!",
				Data:    nil,
				Error:   true,
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
