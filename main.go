package main

import (
    coreConfig "core-payment-lesson/config"
    _ "embed"
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "os"
    "payservice-http-endpoint/config"
    "payservice-http-endpoint/handlers"
    "payservice-http-endpoint/middleware"
)

func main() {
    confPath := os.Args[1]
    conf, err := coreConfig.Parse(confPath)
    if err != nil {
        log.Fatalln(err)
    }
    config.BuildDI(conf)

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

    port := "8080"
    log.Printf("Serving at: http://localhost:%s\n", port)
    log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
