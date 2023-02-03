package transaction

type TransRequest struct {
	DebitAcc  string `json:"debit_acc"`
	CreditAcc string `json:"credit_acc"`
	Amount    int    `json:"amount"`
}

type TransWithDrawRequest struct {
	DebitAcc string `json:"debit_acc"`
	Amount   int    `json:"amount"`
}
type TrnasResponse struct {
	Amount  int    `json:"amount"`
	Message string `json:"message"`
}

type TransactionResp struct {
	Amount       int    `json:"amount"`
	Type         string `json:"type"`
	CreditAcc    string `json:"credit_acc"`
	DebitAcc     string `json:"debit_acc"`
	TransationAt string `json:"transation_at"`
}
type TransactionListResp struct {
	List []TransactionResp `json:"transaction"`
}
