package dto

type ChargeRequest struct {
	Amount      int
	TerminalId  string
	InvoiceId   string
	Description string
}

type ChargeResult struct {
	StatusCode int
	Status     string
	Uuid       string
}
