package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	transQuery     = "INSERT INTO transaction(id,amount,cid_credit,cid_debit,type,transaction_at) VALUES(?,?,?,?,?,?)"
	debit          = "UPDATE customer SET amount = amount-? WHERE id=? "
	credit         = "UPDATE customer SET amount = amount+? WHERE id=?"
	withdraw       = "INSERT INTO transaction(id,amount,cid_debit,type,transaction_at) VALUES(?,?,?,?,?)"
	deposit        = "INSERT INTO transaction(id,amount,cid_credit,type,transaction_at) VALUES(?,?,?,?,?)"
	transactionall = "SELECT * FROM transaction"
	transaction    = "SELECT * FROM transaction where cid_debit=? or cid_credit=?"
)

type TransactionStruct struct {
	Id            string         `db:"id"`
	Amount        int            `db:"amount"`
	CreditAcc     sql.NullString `db:"cid_credit"`
	DebitAcc      sql.NullString `db:"cid_debit"`
	Type          string         `db:"type"`
	TransactionAt string         `db:"transaction_at"`
}

func (s *store) Amounttransaction(ctx context.Context, amount int, creditAcc string, debitAcc string) (err error) {
	return Transaction(ctx, *s.db, &sql.TxOptions{}, func(ctxwithTx context.Context, key string) (err error) {
		txObj := ctxwithTx.Value(key)
		tx, ok := txObj.(*sqlx.Tx)
		if !ok {
			log.Fatalf("error occured while type asserting transaction object from context")
		}

		tx1uuid, err := uuid.NewRandom()
		if err != nil {
			return
		}
		res, err := tx.ExecContext(ctxwithTx, debit, amount, debitAcc)
		if err != nil {
			return err
		}
		count, _ := res.RowsAffected()
		if count == 0 {
			return ErrAccDebitNotExit
		}

		res2, err := tx.ExecContext(ctxwithTx, credit, amount, creditAcc)
		if err != nil {
			return ErrAccCreditNotExit
		}
		count, _ = res2.RowsAffected()
		if count == 0 {
			return ErrAccCreditNotExit
		}
		_, err = tx.ExecContext(ctxwithTx, transQuery, tx1uuid.String(), amount, creditAcc, debitAcc, "txn", time.Now().String())

		return err
	})
}

func (s *store) AmmountWithdraw(ctx context.Context, cid string, amount int) (err error) {
	return Transaction(ctx, *s.db, &sql.TxOptions{}, func(ctxwithTx context.Context, key string) (err error) {
		txObj := ctxwithTx.Value(key)
		tx, ok := txObj.(*sqlx.Tx)
		if !ok {
			log.Fatalf("error occured while type asserting transaction object from context")
		}

		tx1uuid, err := uuid.NewRandom()
		if err != nil {
			return
		}
		res, err := tx.ExecContext(ctxwithTx, debit, amount, cid)
		if err != nil {
			return
		}
		count, _ := res.RowsAffected()
		if count == 0 {
			return ErrAccDebitNotExit
		}

		_, err = tx.ExecContext(ctxwithTx, withdraw, tx1uuid, amount, cid, "withdraw", time.Now().String())
		return
	})
}

func (s *store) AmmountDeposit(ctx context.Context, cid string, amount int) (err error) {
	return Transaction(ctx, *s.db, &sql.TxOptions{}, func(ctxwithTx context.Context, key string) (err error) {
		txObj := ctxwithTx.Value(key)
		tx, ok := txObj.(*sqlx.Tx)
		if !ok {
			log.Fatalf("error occured while type asserting transaction object from context")
		}

		tx1uuid, err := uuid.NewRandom()
		if err != nil {
			return
		}
		res, err := tx.ExecContext(ctxwithTx, credit, amount, cid)
		if err != nil {
			return
		}
		count, _ := res.RowsAffected()
		if count == 0 {
			return ErrAccCreditNotExit
		}

		_, err = tx.ExecContext(ctxwithTx, deposit, tx1uuid, amount, cid, "deposit", time.Now().String())
		return
	})
}

func (s *store) AllTransactionList(ctx context.Context) (list []TransactionStruct, err error) {
	err = s.db.SelectContext(ctx, &list, transactionall)
	return
}
