package handlers

import (
	"diLesson/application/contract"
	"diLesson/application/contract/dto"
	"encoding/json"
	"github.com/golobby/container/v3"
	"log"
	"net/http"
)

type ChargeRequest struct {
	Amount      int    `json:"amount"`
	TerminalId  string `json:"terminalId"`
	InvoiceId   string `json:"invoiceId"`
	Description string `json:"description"`
}

type ChargeResponse struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Uuid       string `json:"uuid"`
}

func Charge(writer http.ResponseWriter, request *http.Request) {

	var charge contract.Charge
	if err := container.Resolve(&charge); err != nil {
		log.Fatalln(err)
	}

	var data ChargeRequest
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := charge.Charge(dto.ChargeRequest{
		Amount:      data.Amount,
		TerminalId:  data.TerminalId,
		InvoiceId:   data.InvoiceId,
		Description: data.Description,
	})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(ChargeResponse{
		StatusCode: result.Status(),
		Status:     result.StatusName(),
		Uuid:       result.Uuid(),
	}); err != nil {
		log.Fatalln(err)
	}
}
