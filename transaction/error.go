package transaction

import "errors"

var (
	ErrInSufficientAmmount = errors.New("amount insufficient")
	ErrUnAuthorize         = errors.New("login user not have access to do this")
)
