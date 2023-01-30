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