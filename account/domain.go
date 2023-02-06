package account

type Account struct {
	UserID  string `json:"user_id"`
	ID      string `json:"id"`
	Balance int    `json:"balance"`
}

type AccountResponse struct {
	AccountInfo Account `json:"balance"`
	Message     string  `json:"message"`
}
type AccountCreate struct {
	UserID string `json:"user_id"`
}
