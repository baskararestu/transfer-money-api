package transactionresponse

type DepositResponse struct {
	Message string         `json:"message"`
	Success bool           `json:"success"`
	Data    CurrentBalance `json:"data"`
}
