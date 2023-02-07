package db

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	transQuery     = "INSERT INTO transaction(id,amount,aid_credit,aid_debit,type,transaction_at) VALUES(?,?,?,?,?,?)"
	debit          = "UPDATE account SET balance = balance-? WHERE id=? "
	credit         = "UPDATE account SET balance = balance+? WHERE id=?"
	withdraw       = "INSERT INTO transaction(id,amount,aid_debit,type,transaction_at) VALUES(?,?,?,?,?)"
	deposit        = "INSERT INTO transaction(id,amount,aid_credit,type,transaction_at) VALUES(?,?,?,?,?)"
	transactionall = "SELECT * FROM transaction"
	transaction    = "SELECT * FROM transaction where aid_debit=? or aid_credit=?"
)

type TransactionStruct struct {
	Id            string         `db:"id"`
	Amount        int            `db:"amount"`
	CreditAcc     sql.NullString `db:"aid_credit"`
	DebitAcc      sql.NullString `db:"aid_debit"`
	Type          string         `db:"type"`
	TransactionAt string         `db:"transaction_at"`
}

func (s *store) Amounttransaction(ctx context.Context, transaction TransactionStruct) (err error) {
	return Transaction(ctx, *s.db, &sql.TxOptions{}, func(ctxwithTx context.Context, key string) (err error) {
		txObj := ctxwithTx.Value(key)
		tx, ok := txObj.(*sqlx.Tx)
		if !ok {
			return errors.New("error occured while type asserting transaction object from context")
		}

		tx1uuid, err := uuid.NewRandom()
		if err != nil {
			return
		}
		parmsListDebit := []interface{}{transaction.Amount, transaction.DebitAcc}

		res, err := tx.ExecContext(ctxwithTx, debit, parmsListDebit...)
		if err != nil {
			return err
		}
		count, err := res.RowsAffected()
		if err != nil {
			return nil
		}
		if count == 0 {
			return ErrAccDebitNotExit
		}

		paramsListCredit := []interface{}{transaction.Amount, transaction.CreditAcc}
		res2, err := tx.ExecContext(ctxwithTx, credit, paramsListCredit...)
		if err != nil {
			return ErrAccCreditNotExit
		}
		count, err = res2.RowsAffected()
		if err != nil {
			return err
		}
		if count == 0 {
			return ErrAccCreditNotExit
		}
		timeCurrent := time.Now().String()
		paramsListtransaction := []interface{}{tx1uuid.String(), transaction.Amount, transaction.CreditAcc, transaction.DebitAcc, "txn", timeCurrent}
		_, err = tx.ExecContext(ctxwithTx, transQuery, paramsListtransaction...)

		return err
	})
}

func (s *store) AmmountWithdraw(ctx context.Context, AccountID string, amount int) (err error) {
	return Transaction(ctx, *s.db, &sql.TxOptions{}, func(ctxwithTx context.Context, key string) (err error) {
		txObj := ctxwithTx.Value(key)
		tx, ok := txObj.(*sqlx.Tx)
		if !ok {
			log.Fatalf("error occured while type asserting transaction object from context")
		}

		txnuuid, err := uuid.NewRandom()
		if err != nil {
			return
		}

		paramsDebit := []interface{}{amount, AccountID}
		res, err := tx.ExecContext(ctxwithTx, debit, paramsDebit...)
		if err != nil {
			return
		}

		count, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if count == 0 {
			return ErrAccDebitNotExit
		}
		timeCurrent := time.Now().String()
		paramsListtransaction := []interface{}{txnuuid, amount, AccountID, "withdraw", timeCurrent}
		_, err = tx.ExecContext(ctxwithTx, withdraw, paramsListtransaction...)
		return
	})
}

func (s *store) AmmountDeposit(ctx context.Context, AccountId string, amount int) (err error) {
	return Transaction(ctx, *s.db, &sql.TxOptions{}, func(ctxwithTx context.Context, key string) (err error) {
		txObj := ctxwithTx.Value(key)
		tx, ok := txObj.(*sqlx.Tx)
		if !ok {
			return errors.New("error occured while type asserting transaction object from context")
		}

		txnuuid, err := uuid.NewRandom()
		if err != nil {
			return
		}

		paramsCredit := []interface{}{amount, AccountId}
		res, err := tx.ExecContext(ctxwithTx, credit, paramsCredit...)
		if err != nil {
			return
		}

		count, err := res.RowsAffected()
		if err != nil {
			return err
		}

		if count == 0 {
			return ErrAccCreditNotExit
		}
		timeCurrent := time.Now().String()
		paramsTransaction := []interface{}{txnuuid, amount, AccountId, "deposit", timeCurrent}
		_, err = tx.ExecContext(ctxwithTx, deposit, paramsTransaction...)
		return
	})
}

func (s *store) AllTransactionList(ctx context.Context) (list []TransactionStruct, err error) {
	err = s.db.SelectContext(ctx, &list, transactionall)
	return
}
