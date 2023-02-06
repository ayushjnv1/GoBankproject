package server

import (
	"github.com/ayushjnv1/Gobank/account"
	"github.com/ayushjnv1/Gobank/app"

	"github.com/ayushjnv1/Gobank/db"
	"github.com/ayushjnv1/Gobank/transaction"
	"github.com/ayushjnv1/Gobank/user"
	userpackage "github.com/ayushjnv1/Gobank/user"
)

type Dependency struct {
	UserService        user.Service
	AccountService     account.Service
	transactionService transaction.Service
}

func initDependency() Dependency {
	appDb := app.GetDB()
	db := db.NewStore(appDb)
	user := userpackage.NewUserService(db)
	account := account.NewAccountService(db)
	transaction := transaction.NewTransactionService(db)

	return Dependency{
		UserService:        user,
		AccountService:     account,
		transactionService: transaction,
	}
}
