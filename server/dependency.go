package server

import (
	"github.com/ayushjnv1/Gobank/account"
	"github.com/ayushjnv1/Gobank/app"
	"github.com/ayushjnv1/Gobank/user"

	"github.com/ayushjnv1/Gobank/db"
	"github.com/ayushjnv1/Gobank/transaction"
)

type Dependency struct {
	UserService        user.Service
	AccountService     account.Service
	transactionService transaction.Service
}

func initDependency() Dependency {
	appDb := app.GetDB()
	db := db.NewStore(appDb)
	user := user.NewUserService(db)
	account := account.NewAccountService(db)
	transaction := transaction.NewTransactionService(db)

	return Dependency{
		UserService:        user,
		AccountService:     account,
		transactionService: transaction,
	}
}
