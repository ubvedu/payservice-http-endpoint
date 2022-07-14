package middleware

import (
    "fmt"
    "github.com/golobby/container/v3"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "payservice-http-endpoint/authentication"
    "payservice-http-endpoint/repository"
)

func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        var users repository.Users
        if err := container.Resolve(&users); err != nil {
            log.Fatalln(err)
        }

        token := mux.Vars(r)["token"]
        w.Header().Set("WWW-Authenticate", fmt.Sprintf("Bearer %s", token))

        id, err := authentication.Id(token)
        if err != nil {
            http.Error(w, err.Error(), http.StatusUnauthorized)
            return
        }

        if !users.HasId(id) {
            http.Error(w, fmt.Sprintf("cannot find user: %s", id.String()), http.StatusForbidden)
            return
        }

        next.ServeHTTP(w, r)
    })
}
