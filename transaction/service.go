package transaction

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ayushjnv1/Gobank/db"
)

type Service interface {
	Amounttransaction(ctx context.Context, transactionRequest TransactionRequest, userID string) (amountRemaining int, err error)
	AmmountWithdraw(ctx context.Context, amount int, debitAcc string) (amountRemaining int, err error)
	AmmountDeposit(ctx context.Context, amount int, creditAcc string) (amountRemaining int, err error)
	AllTransactions(ctx context.Context) (list TransactionListResponse, err error)
}

type transactionService struct {
	db db.Storer
}

func (ts *transactionService) Amounttransaction(ctx context.Context, transactionRequest TransactionRequest, userId string) (amountRemaining int, err error) {
	isSufficient, err := IsAmountSufficient(ctx, transactionRequest.Amount, transactionRequest.DebitAcc, ts.db)
	if err != nil {
		return
	}
	if !isSufficient {
		return amountRemaining, ErrInSufficientAmmount
	}

	isAllow, err := IsLoginUserAccount(ctx, userId, transactionRequest.DebitAcc, ts.db)
	if err != nil {
		return amountRemaining, err
	}
	if !isAllow {
		return amountRemaining, ErrUnAuthorize
	}
	transactionDb := db.TransactionStruct{
		Amount:    transactionRequest.Amount,
		CreditAcc: sql.NullString{String: transactionRequest.CreditAcc, Valid: true},
		DebitAcc:  sql.NullString{String: transactionRequest.DebitAcc, Valid: true},
	}
	err = ts.db.Amounttransaction(ctx, transactionDb)
	if err != nil {
		return
	}

	amountRemaining, err = ts.db.GetAccountBalance(ctx, transactionRequest.DebitAcc)
	return amountRemaining, err
}

func (ts *transactionService) AmmountWithdraw(ctx context.Context, amount int, debitAcc string) (amountRemaining int, err error) {
	isSufficient, err := IsAmountSufficient(ctx, amount, debitAcc, ts.db)
	if err != nil {
		return
	}
	if !isSufficient {
		return amountRemaining, ErrInSufficientAmmount
	}

	err = ts.db.AmmountWithdraw(ctx, debitAcc, amount)
	if err != nil {
		return
	}

	amountRemaining, err = ts.db.GetAccountBalance(ctx, debitAcc)
	return amountRemaining, err
}

func (ts *transactionService) AmmountDeposit(ctx context.Context, amount int, debitAcc string) (amountRemaining int, err error) {
	err = ts.db.AmmountDeposit(ctx, debitAcc, amount)
	if err != nil {
		return
	}
	amountRemaining, err = ts.db.GetAccountBalance(ctx, debitAcc)
	return amountRemaining, err
}

func (ts *transactionService) AllTransactions(ctx context.Context) (list TransactionListResponse, err error) {
	listdbO, err := ts.db.AllTransactionList(ctx)
	if err != nil {
		return
	}

	listTrans := []TransactionResponse{}

	for _, item := range listdbO {
		txn := TransactionResponse{}
		txn.Amount = item.Amount
		txn.CreditAcc = item.CreditAcc.String
		txn.DebitAcc = item.DebitAcc.String
		txn.TransationAt = item.TransactionAt
		txn.Type = item.Type
		listTrans = append(listTrans, txn)
	}
	fmt.Println(listTrans[0].TransationAt)
	list.List = listTrans
	return
}

func NewTransactionService(dbo db.Storer) Service {
	return &transactionService{
		db: dbo,
	}
}
