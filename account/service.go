package account

import (
	"context"

	"github.com/ayushjnv1/Gobank/db"
)

type Service interface {
	CreateAccount(ctx context.Context, uid string) (account Account, err error)
	DeleteAccount(ctx context.Context, id string) (account Account, err error)
	GetAccountBalance(ctx context.Context, id string) (amount int, err error)
}
type AccountService struct {
	store db.Storer
}

func (cs *AccountService) CreateAccount(ctx context.Context, userID string) (account Account, err error) {
	accountCreated, err := cs.store.CreateAccount(ctx, userID)
	if err != nil {
		return account, err
	}
	account = Account{
		Balance: accountCreated.Balance,
		ID:      accountCreated.ID,
		UserID:  accountCreated.UserID,
	}
	return account, err

}
func (cs *AccountService) DeleteAccount(ctx context.Context, accountID string) (account Account, err error) {
	accountDeleted, err := cs.store.DeleteAccount(ctx, accountID)
	if err != nil {
		return account, err
	}
	account = Account{
		Balance: accountDeleted.Balance,
		ID:      accountDeleted.ID,
		UserID:  accountDeleted.UserID,
	}
	return account, err
}

func (cs *AccountService) GetAccountBalance(ctx context.Context, accountID string) (amount int, err error) {
	amount, err = cs.store.GetAccountBalance(ctx, accountID)
	if err != nil {
		return amount, err
	}
	return amount, err
}

func NewAccountService(db db.Storer) Service {
	return &AccountService{
		store: db,
	}
}
