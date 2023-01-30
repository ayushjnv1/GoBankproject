package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ayushjnv1/Gobank/api"
	"github.com/ayushjnv1/Gobank/db"
	"github.com/gorilla/mux"
)


func Login(service Service)(http.HandlerFunc){
	return func(w http.ResponseWriter, r *http.Request) {
		var userLogin UserLogin
		err := json.NewDecoder(r.Body).Decode(&userLogin)
		if err!=nil{
			api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})
			return
		}
		fmt.Println(userLogin.Email)
		user,err:=service.FindByEmail(r.Context(),userLogin.Email)
		if err!=nil{
			if err==db.ErrUserNotExist{
				api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})
				return
			}
            api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})
			return
		}
		if !CheckPasswordHash(userLogin.Password,user.Password){
			api.Error(w,http.StatusBadRequest,api.Response{Message: "Wrong password"})
			return
		}
		jwtString,err := GenerateJWT(user.Id,user.Email,user.Role)
		if(err!=nil){
			api.Error(w,http.StatusInternalServerError,api.Response{Message: err.Error()})
		  return
		}
		api.Success(w,http.StatusAccepted,jwtString)

	}
}






func Create(service Service) (http.HandlerFunc){
	return func(w http.ResponseWriter, r *http.Request) {
		var user UserCreate
		err := json.NewDecoder(r.Body).Decode(&user)
		if err!=nil{
			api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})
			return
		}
		err = service.CreateUser(r.Context(),user)
		if(IsBadRequest(err)){
			api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})	
			return
		}
		if(err!=nil){
			api.Error(w,http.StatusInternalServerError,api.Response{Message: err.Error()})	
			return
		}
		api.Success(w,http.StatusCreated,api.Response{Message: "User Created Succesfully"})

	}
}

func ListOfUser(service Service) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		
		users,err := service.List(r.Context())
		
		if(err==db.ErrUserNotExist){
			api.Error(w,http.StatusNotFound,api.Response{Message: err.Error()})
			return
		}
		if(err!=nil){
          api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})
		  return
		}
		api.Success(w,http.StatusOK,users)
	}
}

func FindById(service Service) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		user,err := service.FindById(r.Context(),vars["id"])
		if(err==db.ErrUserNotExist){
			api.Error(w,http.StatusNotFound,api.Response{Message: err.Error()})
			return
		}
		if(err!=nil){
			api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})
			return		
		}
		api.Success(w,http.StatusOK,user)
	}
}

func DeleteById(service Service) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		vars:= mux.Vars(r)
		
		err := service.DeleteUser(r.Context(),vars["id"]);
		if(err == db.ErrUserNotExist){
			api.Error(w,http.StatusNotFound,api.Response{Message: err.Error()})
			return
		}
		if(err!=nil){
			api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})
			return		
		}
		api.Success(w,http.StatusOK,api.Response{Message:"user Deleted succesfully"})
		
	}
}

func UpdateUserById(service Service) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var user UpdateUser
		err := json.NewDecoder(r.Body).Decode(&user)
		vars := mux.Vars(r)
		id := vars["id"]	
		if err!=nil{
			api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})
			return
		}
		err = service.UpadateUser(r.Context(),user,id)
		
		if(IsBadRequest(err)){
			api.Error(w,http.StatusNotFound,api.Response{Message: err.Error()})
			return	
		}
		if err!=nil{
			api.Error(w,http.StatusBadRequest,api.Response{Message: err.Error()})
			return	
		}
		api.Success(w,http.StatusOK,api.Response{Message:"Updated UserScessfully"})
	}
}