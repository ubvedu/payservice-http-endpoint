package endpoint

import (
    "github.com/gorilla/mux"
    "net/http"
    "payservice-http-endpoint/handlers"
    "payservice-http-endpoint/middleware"
)

func NewRouter() http.Handler {

    router := mux.NewRouter()

    auth := router.Queries("access_token", "{token}").Subrouter()
    auth.Use(middleware.Auth)

    // example:
    //
    //		/charge?access_token=0f62fe
    //
    // must contain json body:
    //
    // 		{
    //		    "amount": 100,
    //		    "terminalId": "x",
    //		    "invoiceId": "y",
    //		    "description": "aboba"
    //		}
    //
    auth.
        HandleFunc("/charge", handlers.Charge).
        Methods(http.MethodPost).
        Headers("Content-Type", "application/json")

    return router
}
