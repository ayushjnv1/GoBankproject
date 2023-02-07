package server

import (
	"net/http"

	"github.com/ayushjnv1/Gobank/account"
	"github.com/ayushjnv1/Gobank/api"
	"github.com/ayushjnv1/Gobank/transaction"
	"github.com/ayushjnv1/Gobank/user"
	"github.com/gorilla/mux"
)

const (
	ADMIN = 1
	USER  = 2
)

func initRouter(dep Dependency) (router *mux.Router) {
	router = mux.NewRouter()

	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	//Login
	router.HandleFunc("/login", user.Login(dep.UserService)).Methods(http.MethodPost)

	// User
	router.HandleFunc("/users", user.Authorize(user.Create(dep.UserService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/users", user.Authorize(user.ListOfUser(dep.UserService), ADMIN)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", user.Authorize(user.FindById(dep.UserService), USER)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", user.Authorize(user.DeleteById(dep.UserService), ADMIN)).Methods(http.MethodDelete)
	router.HandleFunc("/users/{id}", user.Authorize(user.UpdateUserById(dep.UserService), USER)).Methods(http.MethodPut)

	//Account
	router.HandleFunc("/account", user.Authorize(account.CreateAccount(dep.AccountService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/account/{id}", user.Authorize(account.DeleteAccount(dep.AccountService), USER)).Methods(http.MethodDelete)
	router.HandleFunc("/account/{id}/balance", user.Authorize(account.GetAccountBalance(dep.AccountService), USER)).Methods(http.MethodGet)

	//Transaction
	router.HandleFunc("/transaction", user.Authorize(transaction.InitiateTransaction(dep.transactionService), USER)).Methods(http.MethodPost)
	router.HandleFunc("/withdraw", user.Authorize(transaction.WithdrawAmount(dep.transactionService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/deposit", user.Authorize(transaction.DepositAmount(dep.transactionService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/transactions", user.Authorize(transaction.FetchListOfTransaction(dep.transactionService), ADMIN)).Methods(http.MethodGet)
	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "test"})
}
