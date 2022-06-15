package main

import (
	"diLesson/config"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"payservice-http-endpoint/handlers"
	"payservice-http-endpoint/middleware"
)

func main() {
	if err := config.BuildDI(); err != nil {
		log.Panicln(err)
	}

	router := mux.NewRouter()

	auth := router.Queries("token", "{token}").Subrouter()
	auth.Use(middleware.Auth)

	auth.
		HandleFunc("/charge", handlers.Charge).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/json")

	port := "8080"
	log.Printf("Serving at: http://localhost:%s\n", port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
