package transaction

import (
	"context"

	"github.com/ayushjnv1/Gobank/db"
)


type Service interface{
	Amounttransection(ctx context.Context,amount int,creditAcc string,debitAcc string,uid string)(amountRem int,err error)	
	AmmountWithdraw(ctx context.Context,amount int,debitAcc string )(amountRem int ,err error)
	AmmountDeposit(ctx context.Context,amount int,creditAcc string)(amountR int,err error)
}

type transactionService struct{
	db db.Storer
}

func (ts *transactionService)Amounttransection(ctx context.Context,amount int,creditAcc string,debitAcc string,uid string)(amountRem int, err error){
	isSufficient ,err := IsAmountSufficient(ctx,amount,debitAcc,ts.db)   	
	if err!=nil{
		return 
	}
	if !isSufficient{
        return amountRem,ErrInSufficientAmmount
	}  
	
	isAllow,err := IsLoginUserCustomer(ctx,uid,debitAcc,ts.db)
	if err!=nil{
		return amountRem,err
	}
	if !isAllow{
		return amountRem,ErrUnAuthorize
	}

	err = ts.db.Amounttransection(ctx,amount,creditAcc,debitAcc)	
	if err!= nil{
		return 
	}
	amountRem ,err = ts.db.GetammountAcc(ctx,debitAcc)		
	return amountRem,err
}


func( ts *transactionService)AmmountWithdraw(ctx context.Context,amount int,debitAcc string )(amountRem int,err error){
    isSufficient ,err := IsAmountSufficient(ctx,amount,debitAcc,ts.db)   
	if err!=nil{
		return 
	}
	if !isSufficient{
        return amountRem,ErrInSufficientAmmount
	}

	err = ts.db.AmmountWithdraw(ctx,debitAcc,amount)
	if err!=nil{
		return 
	}
	
	amountRem ,err = ts.db.GetammountAcc(ctx,debitAcc)	
	return amountRem,err
}

func (ts *transactionService)AmmountDeposit(ctx context.Context,amount int,debitAcc string)(amountR int,err error){
	err = ts.db.AmmountDeposit(ctx,debitAcc,amount)
	if err!=nil{
		return 
	}
	amountR ,err = ts.db.GetammountAcc(ctx,debitAcc)	
	return amountR,err
}

func NewTransactionService(dbo db.Storer)(Service){
	return &transactionService{
		db:dbo,
	}
}

