package transaction

import (
	"encoding/json"
	"net/http"

	"github.com/ayushjnv1/Gobank/api"
)

func WithdrawAmount(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transReq TransactionRequest
		err := json.NewDecoder(r.Body).Decode(&transReq)

		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		amount, err := service.AmmountWithdraw(r.Context(), transReq.Amount, transReq.DebitAcc)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		api.Success(w, http.StatusAccepted, TrnasactionResponse{Amount: amount, Message: "amount withdraw successful"})
	}
}

func DepositAmount(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transReq TransactionRequest
		err := json.NewDecoder(r.Body).Decode(&transReq)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		amount, err := service.AmmountDeposit(r.Context(), transReq.Amount, transReq.CreditAcc)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		api.Success(w, http.StatusAccepted, TrnasactionResponse{Amount: amount, Message: "amount deposite successful"})
	}
}

func InitiateTransaction(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transactionRequest = TransactionRequest{Amount: -1}
		err := json.NewDecoder(r.Body).Decode(&transactionRequest)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = Validation(transactionRequest)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
		}

		id := r.Header.Get("id")

		amount, err := service.AmountTransaction(r.Context(), transactionRequest, id)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		api.Success(w, http.StatusAccepted, TrnasactionResponse{Amount: amount, Message: "transaction successful"})
	}
}

func FetchListOfTransaction(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := service.AllTransactions(r.Context())
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		api.Success(w, http.StatusOK, response)
	}
}
