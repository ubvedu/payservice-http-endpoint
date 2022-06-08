package handlers

import (
	"encoding/json"
	"github.com/golobby/container/v3"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"payervice-http-endpoint/application/contracts"
	"payervice-http-endpoint/application/contracts/dto"
	"strconv"
)

type ChargeResponse struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Uuid       string `json:"uuid"`
}

func Charge(writer http.ResponseWriter, request *http.Request) {
	var service contracts.PayService
	if err := container.Resolve(&service); err != nil {
		log.Fatalln(err)
	}

	vars := mux.Vars(request)
	amount, err := strconv.Atoi(vars["amount"])
	if err != nil {
		log.Fatalln(err)
	}
	terminalId := vars["terminalId"]
	invoiceId := vars["invoiceId"]
	description, _ := vars["description"]

	result := service.Charge(dto.ChargeRequest{
		Amount:      amount,
		TerminalId:  terminalId,
		InvoiceId:   invoiceId,
		Description: description,
	})

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(ChargeResponse{
		StatusCode: result.StatusCode,
		Status:     result.Status,
		Uuid:       result.Uuid,
	}); err != nil {
		log.Fatalln(err)
	}
}
