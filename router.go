package endpoint

import (
    "github.com/gorilla/mux"
    "net/http"
    "payservice-http-endpoint/handlers"
    "payservice-http-endpoint/middleware"
)

func NewRouter() http.Handler {

    router := mux.NewRouter()

    router.
        HandleFunc("/signIn", handlers.SignIn).
        Methods(http.MethodPost).
        Queries("email", "{email}", "password", "{password}")

    auth := router.Queries("accessToken", "{token}").Subrouter()
    auth.Use(middleware.Auth)

    auth.
        HandleFunc("/charge", handlers.Charge).
        Methods(http.MethodPost).
        Headers("Content-Type", "application/json")

    return router
}
