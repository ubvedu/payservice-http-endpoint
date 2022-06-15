package middleware

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token := mux.Vars(request)["token"]
		// TODO: Authorization
		fmt.Printf("authorized using token: %s\n", token)
		next.ServeHTTP(writer, request)
	})
}
