package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)
const (
	transQuery = "INSERT INTO transaction(id,amount,cid_credit,cid_debit,type,tran_at) VALUES(?,?,?,?,?,SYSDATE())"
	debit = "UPDATE customer SET amount = amount-? WHERE id=? "
	credit = "UPDATE customer SET amount = amount+? WHERE id=?"
	withdraw = "INSERT INTO transaction(id,amount,cid_debit,type,tran_at) VALUES(?,?,?,?,SYSDATE())"
	deposit = "INSERT INTO transaction(id,amount,cid_credit,type,tran_at) VALUES(?,?,?,?,SYSDATE())" 
)

func (s *store) Amounttransection(ctx context.Context,amount int,creditAcc string,debitAcc string)(err error){   
	return Transaction(ctx,*s.db, &sql.TxOptions{},func(ctxwithTx context.Context,key string ) (err error){
      tx := ctxwithTx.Value(key).(*sqlx.Tx);
	  tx1uuid,err :=uuid.NewRandom()	  
	  if(err!=nil){
	  return 
	  }      
	  res,err := tx.ExecContext(ctxwithTx,debit,amount,debitAcc)
	  if(err!=nil){
		return err
	  }
	  count,_ := res.RowsAffected()
	  if count == 0{
		return ErrAccDebitNotExit
	  }

	  res2,err := tx.ExecContext(ctxwithTx,credit,amount,creditAcc)
	  if(err!=nil){
		return ErrAccCreditNotExit
	  } 
	  count,_= res2.RowsAffected()
	  if count == 0{
		return ErrAccCreditNotExit
	  }                
	  _ ,err = tx.ExecContext(ctxwithTx,transQuery,tx1uuid.String(),amount,creditAcc,debitAcc,"txn")
	
	  return err
    })
}

func (s *store) AmmountWithdraw(ctx context.Context,cid string, amount int)(err error){
 return Transaction(ctx,*s.db, &sql.TxOptions{},func(ctxwithTx context.Context, key string) (err error) {
	tx := ctxwithTx.Value(key);
	tx1uuid,err :=uuid.NewRandom()	  
	if(err!=nil){
	return 
	}  
	res,err := tx.(*sqlx.Tx).ExecContext(ctxwithTx,debit,amount,cid)
	if(err!=nil){
		return 
	}
	count,_ := res.RowsAffected()
	if count == 0{
		return ErrAccDebitNotExit
	}
	
	_,err = tx.(*sqlx.Tx).ExecContext(ctxwithTx,withdraw,tx1uuid,amount,cid,"withdraw")
	return 
   })
}

func (s *store) AmmountDeposit(ctx context.Context,cid string, amount int)(err error){
	return Transaction(ctx,*s.db, &sql.TxOptions{},func(ctxwithTx context.Context, key string) (err error) {
	   tx := ctxwithTx.Value(key);
	   tx1uuid,err :=uuid.NewRandom()	  
	   if(err!=nil){
	   return 
	   }  
	   res,err := tx.(*sqlx.Tx).ExecContext(ctxwithTx,credit,amount,cid)
	   if(err!=nil){
		   return 
	   }
	   count,_ := res.RowsAffected()
	   if count == 0{
		   return ErrAccCreditNotExit
	   }
	   
	   _,err = tx.(*sqlx.Tx).ExecContext(ctxwithTx,deposit,tx1uuid,amount,cid,"deposit")
	   return 
	  })
   }
