package handlers

import (
    "core-payment-lesson/server"
    "encoding/json"
    "fmt"
    "github.com/golobby/container/v3"
    "github.com/google/uuid"
    "github.com/liangyaopei/checker"
    "log"
    "net/http"
)

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
    if valid, message := requestData.Check(); !valid {
        http.Error(w, message, http.StatusBadRequest)
        return
    }

    result, err := client.Charge(r.Context(), requestData.Rpc())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(MakeChargeResponse(result)); err != nil {
        log.Fatalln(err)
    }
}

type ChargeRequest struct {
    Amount      int64  `json:"amount"`
    Currency    string `json:"currency"`
    TerminalId  string `json:"terminalId,omitempty"`
    InvoiceId   string `json:"invoiceId,omitempty"`
    Description string `json:"description"`
    CreditCard  struct {
        Number            string `json:"number,omitempty"`
        VerificationValue string `json:"verificationValue,omitempty"`
        Holder            string `json:"holder,omitempty"`
        ExpMonth          int32  `json:"expMonth,omitempty"`
        ExpYear           string `json:"expYear,omitempty"`
        Token             string `json:"token,omitempty"`
    } `json:"creditCard"`
}

func (r *ChargeRequest) Rpc() *server.ChargeRequestMessage {
    return &server.ChargeRequestMessage{
        Amount:      r.Amount,
        Currency:    r.Currency,
        TerminalId:  r.TerminalId,
        InvoiceId:   r.InvoiceId,
        Description: r.Description,
        CreditCard: &server.ChargeRequestMessage_CreditCard{
            Number:                       r.CreditCard.Number,
            VerificationValue:            r.CreditCard.VerificationValue,
            Holder:                       r.CreditCard.Holder,
            ExpMonth:                     server.ChargeRequestMessage_CreditCard_ExpMonth(r.CreditCard.ExpMonth - 1),
            ExpYear:                      r.CreditCard.ExpYear,
            Token:                        r.CreditCard.Token,
            SkipThreeDSecureVerification: false,
        },
    }
}

func (r *ChargeRequest) Check() (valid bool, message string) {
    return checker.And(
        checker.NeInt("Amount", 0),
        UUID("TerminalId"),
        checker.Field("CreditCard", checker.Or(
            checker.And(
                checker.Length("Number", 12, 19),
                checker.Length("VerificationValue", 3, 4),
                checker.NeStr("Holder", ""),
                checker.RangeInt("ExpMonth", 1, 12),
                checker.NeStr("ExpYear", ""),
            ),
            checker.NeStr("Token", ""),
        )),
    ).Check(r)
}

func UUID(fieldExpr string) checker.Rule {
    return checker.Custom(fieldExpr, func(exprValue any) (bool, string) {
        value := exprValue.(string)
        if _, err := uuid.Parse(value); err != nil {
            return false, fmt.Sprintf("%s is not UUID", fieldExpr)
        }
        return true, ""
    })
}

type ChargeResponse struct {
    StatusCode int32  `json:"statusCode"`
    Status     string `json:"status"`
    Uuid       string `json:"uuid"`
}

func MakeChargeResponse(rpc *server.ChargeResponseMessage) ChargeResponse {
    return ChargeResponse{
        StatusCode: rpc.GetStatusCode(),
        Status:     rpc.GetStatusName(),
        Uuid:       rpc.GetUuid(),
    }
}
