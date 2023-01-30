package user

import (
	"context"

	"github.com/ayushjnv1/Gobank/db"
)

type Service interface{
	List(ctx context.Context)(user UserList,err error)
	FindById(ctx context.Context,id string)(user UserResponse,err error)
	FindByEmail	(ctx context.Context,email string)(user UserCreate,err error)
	UpadateUser(ctx context.Context,user UpdateUser,id string)(err error)
	DeleteUser(ctx context.Context,id string)(err error)
	CreateUser(ctx context.Context,user UserCreate)(err error)
}

type userService struct{
	store db.Storer
}

func (us *userService)UpadateUser(ctx context.Context,user UpdateUser,id string)(err error){
   err = UpdateUserValidate(user)
   if err!=nil {
    return
   }
   err = us.store.UpdateUser(ctx,db.User{
	Name:user.Name,
	Email: user.Email,
	Role: user.Role,
},id )
   return 
}

func (us *userService) List( ctx context.Context)(resp UserList,err error){
	userDbObj,err := us.store.ListOfUser(ctx)	
	if err!=nil{
		return 
	}

   userlist := []UserResp{}
	for _, u := range(userDbObj){
		ur := UserResp{}
		ur.Email = u.Email
		ur.Name = u.Name
		ur.Role = u.Role
		ur.Id = u.Id
		userlist = append(userlist,ur);
	}
	
	resp.User = userlist
	
	return 
}

func (us *userService) CreateUser(ctx context.Context,user UserCreate)(err error){
  err = ValidateCreateUser(user)
  if(err!=nil){
	return
  }
  password,err:= HashPassword(user.Password)
	if err!=nil{
		return
	}
  err =  us.store.CreateUser(ctx,db.User{
	Name:user.Name,
	Email: user.Email,
	Role: user.Role,
	Password: password,
  })
  return 
}

func (us *userService) DeleteUser(ctx context.Context,id string) (err error){
   err = us.store.DeleteUser(ctx,id)  
   return
}

func (us *userService) FindById(ctx context.Context,id string) (user UserResponse,err error){
userdbo,err := us.store.FindById(ctx,id)
if(err!=nil){
	return
}
user.User=UserResp{
	Email: userdbo.Email,
	Name: userdbo.Name,
	Role:  userdbo.Role,
	Id: userdbo.Id,
	
}
return
}

func (us *userService) FindByEmail(ctx context.Context,email string) (user UserCreate,err error){
	userdbo,err := us.store.FindByEmail(ctx,email)
	if(err!=nil){
		return
	}
	user =UserCreate{
		Email: userdbo.Email,
		Name: userdbo.Name,
		Role:  userdbo.Role,
		Id: userdbo.Id,		
		Password: userdbo.Password,
	}
	return
	}


func NewUserService(db db.Storer) (Service){
 return &userService{
	store: db,
}
}