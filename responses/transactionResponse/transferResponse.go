package transactionresponse

type TransferResponse struct {
	Message string       `json:"message"`
	Success bool         `json:"success"`
	Data    TransferData `json:"data,omitempty"`
}

type TransferData struct {
	Amount       float64 `json:"amount"`
	FromAccount  string  `json:"from_account"`
	SenderName   string  `json:"sender_name"`
	ToAccount    string  `json:"to_account"`
	ReceiverName string  `json:"receiver_name"`
}

type TransferServiceResult struct {
	Success bool
	Message string
	Data    TransferData
	Error   error
}
