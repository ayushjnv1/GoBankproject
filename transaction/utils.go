package transaction

import (
	"context"

	"github.com/ayushjnv1/Gobank/db"
)
func  IsAmountSufficient(ctx context.Context, amount int,acc string,db db.Storer)(f bool,err error){
	amm,err := db.GetammountAcc(ctx,acc)
	if err!=nil{
		return false,err
	}
	if amm<amount{
		return false,err
	}
	return true,err
}

func IsLoginUserCustomer(ctx context.Context,uid string, cutid string,db db.Storer)(f bool,err error){
   cust,err := db.GetCustomer(ctx,cutid)
   if err!=nil{
	return false,err
   }
   if(cust.Uid!=uid){
	return false,err
   }
   return true,err
}