package customer



type Customer struct{
	Uid string `json:"uid"`
	Id string  `json:"id"`
	Amount int `json:"amount"`
}
type CustomerRes struct{
	Cus Customer `json:"cus"`
	Message string `josn:"message"`
}
type CutomerCreate struct{	
	Uid string `json:"uid"`	
}

func ValidateCustomer(cus CutomerCreate){

}

