package transaction



type TransRequest struct{
	DebitAcc string `json:"debit_acc"`
	CreditAcc string `json:"credit_acc"`
	Amount int `json:"amount"`
}

type TransWithDrawRequest struct{
	DebitAcc string `json:"debit_acc"`
	Amount int `json:"amount"`
}
type TrnasResponse struct{
   Amount int `json:"amount"`
   Message string `json:"message"`
}

