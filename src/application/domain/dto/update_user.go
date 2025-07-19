package dto

type UpdateUserReq struct {
	UserName           string `json:"username"`
	LastName           string `json:"lastName"`
	FirstName          string `json:"firstName"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	Phone              string `json:"phone"`
	EmailNotifications bool   `json:"emailNotifications"`
	SmsNotifications   bool   `json:"smsNotifications"`
	Language           string `json:"language"`
}

type UpdateUserRes struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"username"`
	Phone    string `json:"phone"`
	Language string `json:"language"`
}
