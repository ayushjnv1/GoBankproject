package server

import (
	"github.com/ayushjnv1/Gobank/app"
	"github.com/ayushjnv1/Gobank/customer"
	"github.com/ayushjnv1/Gobank/db"
	"github.com/ayushjnv1/Gobank/transaction"
	"github.com/ayushjnv1/Gobank/user"
)


type Dependency struct {
	UserService user.Service
	CustomerService customer.Service
	transactionService transaction.Service
}

func initDependency()(Dependency){
	appDb := app.GetDB()
	db:= db.NewStore(appDb)
	user := user.NewUserService(db)
	customer := customer.NewCustomerService(db)
	transaction:= transaction.NewTransactionService(db)
	return Dependency{
		UserService: user,
		CustomerService: customer,
		transactionService: transaction,
		};
}