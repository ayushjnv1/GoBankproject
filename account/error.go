package account

import "errors"

var (
	ErrUidExit = errors.New("user of this uid is not present")
)
