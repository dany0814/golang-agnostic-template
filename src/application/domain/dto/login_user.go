package dto

type LoginUserReq struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRes struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"username"`
	Phone    string `json:"phone"`
	Language string `json:"language"`
	Token    string `json:"token"`
}
