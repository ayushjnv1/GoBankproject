package transaction

import (
	"context"
	"fmt"

	"github.com/ayushjnv1/Gobank/db"
)

type Service interface {
	Amounttransaction(ctx context.Context, amount int, creditAcc string, debitAcc string, uid string) (amountRem int, err error)
	AmmountWithdraw(ctx context.Context, amount int, debitAcc string) (amountRem int, err error)
	AmmountDeposit(ctx context.Context, amount int, creditAcc string) (amountR int, err error)
	AllTransactionList(ctx context.Context) (list TransactionListResp, err error)
}

type transactionService struct {
	db db.Storer
}

func (ts *transactionService) Amounttransaction(ctx context.Context, amount int, creditAcc string, debitAcc string, uid string) (amountRem int, err error) {
	isSufficient, err := IsAmountSufficient(ctx, amount, debitAcc, ts.db)
	if err != nil {
		return
	}
	if !isSufficient {
		return amountRem, ErrInSufficientAmmount
	}

	isAllow, err := IsLoginUserCustomer(ctx, uid, debitAcc, ts.db)
	if err != nil {
		return amountRem, err
	}
	if !isAllow {
		return amountRem, ErrUnAuthorize
	}

	err = ts.db.Amounttransaction(ctx, amount, creditAcc, debitAcc)
	if err != nil {
		return
	}
	amountRem, err = ts.db.GetAccountBalance(ctx, debitAcc)
	return amountRem, err
}

func (ts *transactionService) AmmountWithdraw(ctx context.Context, amount int, debitAcc string) (amountRem int, err error) {
	isSufficient, err := IsAmountSufficient(ctx, amount, debitAcc, ts.db)
	if err != nil {
		return
	}
	if !isSufficient {
		return amountRem, ErrInSufficientAmmount
	}

	err = ts.db.AmmountWithdraw(ctx, debitAcc, amount)
	if err != nil {
		return
	}

	amountRem, err = ts.db.GetAccountBalance(ctx, debitAcc)
	return amountRem, err
}

func (ts *transactionService) AmmountDeposit(ctx context.Context, amount int, debitAcc string) (amountR int, err error) {
	err = ts.db.AmmountDeposit(ctx, debitAcc, amount)
	if err != nil {
		return
	}
	amountR, err = ts.db.GetAccountBalance(ctx, debitAcc)
	return amountR, err
}

func (ts *transactionService) AllTransactionList(ctx context.Context) (list TransactionListResp, err error) {
	listdbO, err := ts.db.AllTransactionList(ctx)
	if err != nil {
		return
	}

	listTrans := []TransactionResp{}

	for _, item := range listdbO {
		txn := TransactionResp{}
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
