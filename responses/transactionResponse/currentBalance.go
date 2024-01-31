package transactionresponse

type CurrentBalance struct {
	AccountName    string  `json:"account_name"`
	AccountNumber  string  `json:"account_number"`
	CurrentBalance float64 `json:"current_balance"`
}
