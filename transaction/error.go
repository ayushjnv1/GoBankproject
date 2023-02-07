package transaction

import "errors"

var (
	ErrInSufficientAmmount     = errors.New("amount insufficient")
	ErrUnAuthorize             = errors.New("login user not have access to do this")
	ErrDebitAccountNotExist    = errors.New("debit account details not present ")
	ErrorCreditAccountNotExist = errors.New("credit account details not present")
	ErrorAmountNotExist        = errors.New("amount value not present")
)
