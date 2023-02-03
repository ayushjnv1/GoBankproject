package customer

import (
	"encoding/json"
	"net/http"

	"github.com/ayushjnv1/Gobank/api"
	"github.com/gorilla/mux"
)

func CreateCustomer(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cust CutomerCreate
		err := json.NewDecoder(r.Body).Decode(&cust)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		cus, err := service.CreateCustomer(r.Context(), cust.Userid)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		api.Success(w, http.StatusCreated, CustomerRes{Message: "customer createded successfully", CustomerInfo: cus})

	}
}
func DeleteCustomer(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		cid := vars["id"]
		cus, err := service.DeleteCustomer(r.Context(), cid)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		api.Success(w, http.StatusOK, CustomerRes{Message: "customer deleted successfully", CustomerInfo: cus})
	}
}

func GetammountAcc(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cid := vars["id"]
		ammount, err := service.GetAccountBalance(r.Context(), cid)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		api.Success(w, http.StatusOK, struct {
			Amount int `json:"amount"`
		}{Amount: ammount})
	}
}
