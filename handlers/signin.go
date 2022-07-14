package handlers

import (
    "encoding/json"
    "github.com/golobby/container/v3"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "payservice-http-endpoint/authentication"
    "payservice-http-endpoint/repository"
)

func SignIn(w http.ResponseWriter, r *http.Request) {

    var users repository.Users
    if err := container.Resolve(&users); err != nil {
        log.Fatalln(err)
    }

    vars := mux.Vars(r)
    email := vars["email"]
    password := vars["password"]

    token, err := authentication.Sign(users.AddUser(email, password))
    if err != nil {
        log.Fatalln(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(SignInResponse{AccessToken: token}); err != nil {
        log.Fatalln(err)
    }
}

type SignInResponse struct {
    AccessToken string `json:"accessToken"`
}
