package customer

import (
	"context"

	"github.com/ayushjnv1/Gobank/db"
)


type  Service interface{
	CreateCustomer(ctx context.Context,uid string) (cust Customer,err error)
	DeleteCustomer(ctx context.Context,id string) (cust Customer,err error)
	AmountGet(ctx context.Context,id string)(amount int,err error)

}
type customerService struct{
	store db.Storer
}

func (cs *customerService)CreateCustomer(ctx context.Context,uid string)(cust Customer,err error){
   cus,err := cs.store.CreateCustomer(ctx, uid)
   if err!=nil{
	return cust,err
   }
   cust.Amount = cus.Amount
   cust.Id = cus.Id
   cust.Uid = cus.Uid
   return cust,err

}
func (cs *customerService) DeleteCustomer(ctx context.Context,cid string)(cust Customer ,err error){
cus,err := cs.store.DeleteCustomer(ctx,cid)
if err!=nil{
	return cust,err
}
cust.Amount = cus.Amount
cust.Id = cus.Id
cust.Uid = cus.Uid
return cust,err
}

func (cs *customerService)AmountGet(ctx context.Context,cid string) (amount int,err error){
amount,err = cs.store.GetammountAcc(ctx,cid)
if err!=nil{
	return amount,err
}
return amount,err
}

func NewCustomerService(db db.Storer)(Service){
	return &customerService{
		store: db,
	}
}