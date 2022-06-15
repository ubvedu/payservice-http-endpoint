package handlers

import (
    "encoding/json"
    "github.com/golobby/container/v3"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "payservice-core/application/contract"
    "payservice-core/application/contract/dto"
    "strconv"
)

func Charge(writer http.ResponseWriter, request *http.Request) {
    var charge contract.Charge
    if err := container.Resolve(&charge); err != nil {
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

    result, err := charge.Charge(dto.ChargeRequest{
        Amount:      amount,
        TerminalId:  terminalId,
        InvoiceId:   invoiceId,
        Description: description,
    })
    if err != nil {
        http.Error(writer, err.Error(), http.StatusBadRequest)
    }

    writer.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(writer).Encode(ChargeResponse{
        StatusCode: result.Status(),
        Status:     result.StatusName(),
        Uuid:       result.Uuid(),
    }); err != nil {
        log.Fatalln(err)
    }
}

type ChargeResponse struct {
    StatusCode int    `json:"statusCode"`
    Status     string `json:"status"`
    Uuid       string `json:"uuid"`
}
