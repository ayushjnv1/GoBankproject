package transaction

type TransactionRequest struct {
	DebitAcc  string `json:"debit_acc"`
	CreditAcc string `json:"credit_acc"`
	Amount    int    `json:"amount"`
}

type TransactionWithDrawRequest struct {
	DebitAcc string `json:"debit_acc"`
	Amount   int    `json:"amount"`
}
type TrnasactionResponse struct {
	Amount  int    `json:"amount"`
	Message string `json:"message"`
}

type TransactionResponse struct {
	Amount       int    `json:"amount"`
	Type         string `json:"type"`
	CreditAcc    string `json:"credit_acc"`
	DebitAcc     string `json:"debit_acc"`
	TransationAt string `json:"transation_at"`
}
type TransactionListResponse struct {
	List []TransactionResponse `json:"transaction"`
}

func Validation(txn TransactionRequest) error {
	if txn.Amount < 0 {
		return ErrorAmountNotExist
	}
	if txn.CreditAcc == "" {
		return ErrorCreditAccountNotExist
	}
	if txn.DebitAcc == "" {
		return ErrDebitAccountNotExist
	}
	return nil

}
