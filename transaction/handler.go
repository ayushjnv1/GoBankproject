package transaction

import (
	"encoding/json"
	"net/http"

	"github.com/ayushjnv1/Gobank/api"
)

func AmountWithdraw(service Service) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var transReq TransRequest
		err := json.NewDecoder(r.Body).Decode(&transReq)
		// ask it is possible that we can empty any field & if possible how ?
		if err!=nil{
			api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})
			return
		}
		amount,err:=service.AmmountWithdraw(r.Context(),transReq.Amount,transReq.DebitAcc)
        if err!=nil{
			api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})
			return
		}
		api.Success(w,http.StatusAccepted,TrnasResponse{Amount: amount,Message: "amount withdraw successful"})
	}
}

func AmmountDeposit(service Service) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var transReq TransRequest
		err := json.NewDecoder(r.Body).Decode(&transReq)
		// ask it is possible that we can empty any field & if possible how ?
		if err!=nil{
			api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})
			return
		}
		amount,err:=service.AmmountDeposit(r.Context(),transReq.Amount,transReq.CreditAcc)
        if err!=nil{
			api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})
			return
		}
		api.Success(w,http.StatusAccepted,TrnasResponse{Amount: amount,Message: "amount deposite successful"})
	}
}

func Amounttransection(service Service) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var transReq TransRequest
		err := json.NewDecoder(r.Body).Decode(&transReq)
		if err!=nil{
			api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})
			return
		}
		id := r.Header.Get("id")
		amount,err := service.Amounttransection(r.Context(),transReq.Amount,transReq.CreditAcc,transReq.DebitAcc,id)
		if err!=nil{
			api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})
			return
		}
		api.Success(w,http.StatusAccepted,TrnasResponse{Amount: amount,Message: "transaction successful"})
	
	}
}