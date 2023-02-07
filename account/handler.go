package account

import (
	"encoding/json"
	"net/http"

	"github.com/ayushjnv1/Gobank/api"
	"github.com/gorilla/mux"
)

func CreateAccount(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var account AccountCreate
		err := json.NewDecoder(r.Body).Decode(&account)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		cus, err := service.CreateAccount(r.Context(), account.UserID)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		api.Success(w, http.StatusCreated, AccountResponse{Message: "account createded successfully", AccountInfo: cus})

	}
}

// DeleteAccount
func DeleteAccount(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		accountID := vars["id"]
		account, err := service.DeleteAccount(r.Context(), accountID)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		api.Success(w, http.StatusOK, AccountResponse{Message: "account deleted successfully", AccountInfo: account})
	}
}

// ayush: camel case naming, GetAccountBalance
func GetAccountBalance(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		accountID := vars["id"]
		ammount, err := service.GetAccountBalance(r.Context(), accountID)
		if err != nil {
			api.Error(w, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		api.Success(w, http.StatusOK, struct {
			Amount int `json:"balance"`
		}{Amount: ammount})
	}
}
