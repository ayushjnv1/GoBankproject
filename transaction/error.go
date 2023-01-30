package transaction

import "errors"

var (
	ErrInSufficientAmmount = errors.New("Amount insufficient") 
	ErrUnAuthorize = errors.New("Login user not have access to do this")
)