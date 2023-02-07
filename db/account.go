package db

import (
	"context"

	"github.com/google/uuid"
)

const (
	createAccount = `INSERT INTO account (id,balance,user_id,created_at) VALUES(?,?,?,SYSDATE())`
	deleteAccount = `UPDATE account SET deleted_at=SYSDATE() WHERE id =? AND deleted_at IS NULL`
	getBalance    = `SELECT balance FROM account WHERE id=?`
	getAccount    = `SELECT id,user_id,balance FROM account WHERE id = ?`
)

type Account struct {
	UserID  string `db:"user_id"`
	Balance int    `db:"balance"`
	ID      string `db:"id"`
}

func (s *store) GetAccountBalance(ctx context.Context, accountID string) (balance int, err error) {
	balanceList := []int{}
	err = s.db.SelectContext(ctx, &balanceList, getBalance, accountID)
	if err != nil {
		return
	}

	if len(balanceList) == 0 {
		return 0, ErrUserNotExist
	}

	return balanceList[0], err
}

func (s *store) CreateAccount(ctx context.Context, userId string) (account Account, err error) {
	AccountID, err := uuid.NewUUID()
	if err != nil {
		return account, err
	}
	params := []interface{}{AccountID.String(), 0, userId}
	_, err = s.db.ExecContext(ctx, createAccount, params...)
	if err != nil {
		return account, err
	}
	account, err = s.GetAccount(ctx, AccountID.String())
	if err != nil {
		return account, err
	}
	return account, err
}

func (s *store) DeleteAccount(ctx context.Context, accountID string) (account Account, err error) {
	res, err := s.db.ExecContext(ctx, deleteAccount, accountID)
	if err != nil {
		return account, err
	}
	val, err := res.RowsAffected()
	if err != nil {
		return account, err
	}

	if val == 0 {
		return account, ErrUserNotExist
	}
	account, err = s.GetAccount(ctx, accountID)
	if err != nil {
		return account, err
	}
	return account, err
}

func (s *store) GetAccount(ctx context.Context, AccountID string) (account Account, err error) {
	accountList := []Account{}
	err = s.db.SelectContext(ctx, &accountList, getAccount, AccountID)
	if err != nil {
		return account, err
	}
	account = accountList[0]

	return account, err
}
