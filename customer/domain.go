package customer

type Customer struct {
	Userid  string `json:"uid"`
	ID      string `json:"id"`
	Balance int    `json:"amount"`
}
type CustomerRes struct {
	CustomerInfo Customer `json:"cus"`
	Message      string   `json:"message"`
}
type CutomerCreate struct {
	Userid string `json:"uid"`
}
