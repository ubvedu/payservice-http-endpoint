package middleware

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)

func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := mux.Vars(r)["token"]
        // TODO: Authorization
        fmt.Printf("authorized using token: %s\n", token)
        next.ServeHTTP(w, r)
    })
}
