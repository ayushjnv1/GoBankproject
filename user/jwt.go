package user

import (
	"net/http"
	"time"

	"github.com/ayushjnv1/Gobank/api"
	"github.com/golang-jwt/jwt"
)

type JWTClaim struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

var jwtKey = []byte("yu78jhe5$r")
var role_map = map[string]int{"user": 2, "admin": 1,}

func Authorize(handler http.HandlerFunc, role int ) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
		claims := JWTClaim{}

		if tokenString!=""{
		 token,err :=  new(jwt.Parser).ParseWithClaims(tokenString,&claims,func(t *jwt.Token) (interface{}, error) {return jwtKey,nil})
        if err!=nil{			
				api.Error(w,http.StatusUnauthorized,api.Response{Message: err.Error()})
				return 
			}
		if !token.Valid{
			api.Error(w,http.StatusUnauthorized,api.Response{Message: err.Error()})
		}
			
		 if role_map[claims.Role]>role{
			api.Error(w,http.StatusUnauthorized,api.Response{Message: "You dont have access to perform this action"})	 	
			return
		 }
		 r.Header.Set("id",claims.ID);
	     handler.ServeHTTP(w,r)
		return
		}
		api.Success(w,http.StatusUnauthorized,api.Response{Message: "Please enter token string"})		
	}
}

func GenerateJWT(ID string, Email string, Role string) (tokenString string, err error) {

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &JWTClaim{
		ID:    ID,
		Email: Email,
		Role:  Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}