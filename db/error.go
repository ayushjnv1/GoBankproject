package db

import "errors"

var (
	ErrUserNotExist     = errors.New("User not Exist")
	ErrAccCreditNotExit = errors.New("credit Account not exist ")
	ErrAccDebitNotExit  = errors.New("debit Account not exist")
)
