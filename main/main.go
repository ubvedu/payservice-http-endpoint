package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"http-endpoint/config"
	"http-endpoint/handlers"
	"log"
	"net/http"
)

func main() {
	if err := config.BuildDI(); err != nil {
		log.Panicln(err)
	}

	router := mux.NewRouter()

	// example:
	// 		/charge?amount=10000&terminalId=foo&invoiceId=bar&description=aboba
	router.
		Path("/charge").
		Queries(
			"amount", "{amount:[0-9]+}",
			"terminalId", "{terminalId}",
			"invoiceId", "{invoiceId}",
			"description", "{description}",
		).
		HandlerFunc(handlers.Charge)

	port := "8080"
	log.Printf("Serving at: http://localhost:%s\n", port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
