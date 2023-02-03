package customer

import (
	"context"

	"github.com/ayushjnv1/Gobank/db"
)

type Service interface {
	CreateCustomer(ctx context.Context, uid string) (cust Customer, err error)
	DeleteCustomer(ctx context.Context, id string) (cust Customer, err error)
	GetAccountBalance(ctx context.Context, id string) (amount int, err error)
}
type customerService struct {
	store db.Storer
}

func (cs *customerService) CreateCustomer(ctx context.Context, uid string) (customer Customer, err error) {
	cus, err := cs.store.CreateCustomer(ctx, uid)
	if err != nil {
		return customer, err
	}
	customer = Customer{
		Balance: cus.Balance,
		ID:      cus.ID,
		Userid:  cus.Userid,
	}
	return customer, err

}
func (cs *customerService) DeleteCustomer(ctx context.Context, cid string) (customer Customer, err error) {
	cus, err := cs.store.DeleteCustomer(ctx, cid)
	if err != nil {
		return customer, err
	}
	customer = Customer{
		Balance: cus.Balance,
		ID:      cus.ID,
		Userid:  cus.Userid,
	}
	return customer, err
}

func (cs *customerService) GetAccountBalance(ctx context.Context, cid string) (amount int, err error) {
	amount, err = cs.store.GetAccountBalance(ctx, cid)
	if err != nil {
		return amount, err
	}
	return amount, err
}

func NewCustomerService(db db.Storer) Service {
	return &customerService{
		store: db,
	}
}
