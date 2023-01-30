package server

import (
	"net/http"

	"github.com/ayushjnv1/Gobank/api"
	"github.com/ayushjnv1/Gobank/customer"
	"github.com/ayushjnv1/Gobank/transaction"
	"github.com/ayushjnv1/Gobank/user"
	"github.com/gorilla/mux"
)

func initRouter(dep Dependency)(router *mux.Router){
	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	//Login
	router.HandleFunc("/login",user.Login(dep.UserService)).Methods(http.MethodPost)	

	// User
	router.HandleFunc("/users",user.Authorize(user.Create(dep.UserService),1)).Methods(http.MethodPost)
	router.HandleFunc("/users",user.Authorize(user.ListOfUser(dep.UserService),1)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}",user.Authorize(user.FindById(dep.UserService),2)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}",user.Authorize(user.DeleteById(dep.UserService),1)).Methods(http.MethodDelete)
	router.HandleFunc("/users/{id}",user.Authorize(user.UpdateUserById(dep.UserService),2)).Methods(http.MethodPut)

	//Customer
	router.HandleFunc("/customer",user.Authorize(customer.CreateCustomer(dep.CustomerService),1)).Methods(http.MethodPost)
	router.HandleFunc("/customer/{id}",user.Authorize(customer.DeleteCustomer(dep.CustomerService),2)).Methods(http.MethodDelete)
	router.HandleFunc("/customer/amount/{id}",user.Authorize(customer.GetammountAcc(dep.CustomerService),2)).Methods(http.MethodGet)

	//transaction
	router.HandleFunc("/amountTransaction",user.Authorize(transaction.Amounttransection(dep.transactionService),2)).Methods(http.MethodPost)
	router.HandleFunc("/amountWithdraw",user.Authorize(transaction.AmountWithdraw(dep.transactionService),1)).Methods(http.MethodPost)
	router.HandleFunc("/amountDeposit",user.Authorize(transaction.AmmountDeposit(dep.transactionService),1)).Methods(http.MethodPost)
	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}