package db

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Storer interface{
	// User
	CreateUser(ctx context.Context,user User) (err error)
	DeleteUser(ctx context.Context,id string)(err error)
	UpdateUser(ctx context.Context,user User,id string) (err error)
	UpdatePassword(ctx context.Context,pass string,Id string)(err error)
	ListOfUser(ctx context.Context) (user []User,err error)
	FindById(ctx context.Context,id string)(user User,err error)
	FindByEmail(ctx context.Context,email string)(user User,err error)

	//transection
	AmmountWithdraw(ctx context.Context,cid string, amount int)(err error)
	AmmountDeposit(ctx context.Context,cid string, amount int)(err error)
	Amounttransection(ctx context.Context,amount int,creditAcc string,debitAcc string)(err error)

	//customer
	GetammountAcc(ctx context.Context,id string)(amount int, err error)
	CreateCustomer(ctx context.Context,uid string)(cust Customer,err error)
	GetCustomer(ctx context.Context,cid string)(cust Customer,err error)
	DeleteCustomer(ctx context.Context,cid string)(cust Customer,err error)

} 
type store struct{
	db *sqlx.DB
  }


  func Transaction(ctx context.Context,dbx sqlx.DB,opt *sql.TxOptions,fun func(context.Context,string)(error))(err error){
	tx,err :=  dbx.BeginTxx(ctx,opt);
	if(err!=nil){
		return
	}
	defer func() {
		if p := recover(); p != nil {
			switch p:=p.(type){
			case error:
				err = errors.WithStack(p)
		 	default:
				err= errors.Errorf("s",p)	
			}
		}
	
	    if err !=nil {
             e:= tx.Rollback()
		     if e!= nil{
			err=errors.WithStack(e)
		    }
		 return
	}    
	err = errors.WithStack(tx.Commit());
	}()

	ctxwithTx := newContext(ctx,tx,"txOb")
	err = fun(ctxwithTx,"txOb")
	return err
  }

  func newContext(ctx context.Context,val interface{},key string)context.Context{
    return context.WithValue(ctx,key,val)
  }

  
  func NewStore(dbo *sqlx.DB)(Storer){
    return &store{
		db:dbo,
	}
  }