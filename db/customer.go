package db

import (
	"context"

	"github.com/google/uuid"
)

const(
	createCustomer = `INSERT INTO customer (id,amount,uid,created_at) VALUES(?,?,?,SYSDATE())`
	deleteCustomer = `UPDATE customer SET deleted_at=SYSDATE() WHERE id =? AND deleted_at IS NULL`
	getAmount = `SELECT amount FROM customer WHERE id=?`
	getCustomer = `SELECT id,uid,amount FROM customer WHERE id = ?`	
)

type Customer struct{
	Uid string `db:"uid"`
	Amount int `db:"amount"`
	Id string  `db:"id"`

}


func (s *store)GetammountAcc(ctx context.Context, cid string)(amount int,err error){
    amo := []int{}	
	err = s.db.SelectContext(ctx,&amo,getAmount,cid)	
   if(err!=nil){
	return 
   }
   if len(amo)==0{
	return 0,ErrUserNotExist
   }
   
   return amo[0],err
}

func (s *store)CreateCustomer(ctx context.Context,uid string)(cust Customer, err error){
	uuid1,err := uuid.NewUUID()
	if(err!=nil){
		return cust,err
	}
	_,err = s.db.ExecContext(ctx,createCustomer,uuid1.String(),0,uid,)
	if err!= nil{
		return cust,err
	}
	cust,err = s.GetCustomer(ctx,uuid1.String())
	if err!=nil{
		return cust,err
	}
	return cust,err	
}

func (s *store)DeleteCustomer(ctx context.Context,cid string)(cust Customer,err error){
	res,err := s.db.ExecContext(ctx,deleteCustomer,cid);
	if(err!=nil){
		return cust,err
	}
	val,_:= res.RowsAffected()

	if(val==0){
		return cust,ErrUserNotExist
	}
	cust,err = s.GetCustomer(ctx,cid)
	if err!=nil{
		return cust,err
	}
	return cust, err 
}

func (s *store)GetCustomer(ctx context.Context, id string)(cust Customer,err error){
	a:= []Customer{}
	err = s.db.SelectContext(ctx,&a,getCustomer,id);
	if err!=nil{
		return cust,err
	}
	cust = a[0]

	return cust,err
}