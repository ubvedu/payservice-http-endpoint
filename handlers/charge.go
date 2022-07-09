package handlers

import (
    "core-payment-lesson/server"
    "encoding/json"
    "github.com/golobby/container/v3"
    "github.com/google/uuid"
    "log"
    "net/http"
)

type ChargeRequest struct {
    Amount      int64  `json:"amount"`
    TerminalId  string `json:"terminalId"`
    InvoiceId   string `json:"invoiceId"`
    Description string `json:"description"`
}

type ChargeResponse struct {
    StatusCode int32  `json:"statusCode"`
    Status     string `json:"status"`
    Uuid       string `json:"uuid"`
}

func Charge(w http.ResponseWriter, r *http.Request) {

    var client server.PayServiceClient
    if err := container.Resolve(&client); err != nil {
        log.Fatalln(err)
    }

    var requestData ChargeRequest
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    rpcRequestData := requestData.Rpc()
    log.Println("new request:", rpcRequestData)

    result, err := client.Charge(r.Context(), requestData.Rpc())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(NewChargeResponse(result)); err != nil {
        log.Fatalln(err)
    }
}

func (r *ChargeRequest) Rpc() *server.ChargeRequestMessage {
    return &server.ChargeRequestMessage{
        Amount:      r.Amount,
        Currency:    "RUB",
        TerminalId:  uuid.New().String(),
        InvoiceId:   uuid.New().String(),
        Description: r.Description,
        CreditCard: &server.ChargeRequestMessage_CreditCard{
            Number:                       "0123456789101112",
            VerificationValue:            "123",
            Holder:                       "MOMENTUM R",
            ExpMonth:                     10,
            ExpYear:                      "2025",
            Token:                        "",
            SkipThreeDSecureVerification: false,
        },
    }
}

func NewChargeResponse(rpc *server.ChargeResponseMessage) *ChargeResponse {
    return &ChargeResponse{
        StatusCode: rpc.GetStatusCode(),
        Status:     rpc.GetStatusName(),
        Uuid:       rpc.GetUuid(),
    }
}
