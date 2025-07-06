package dto

type RegisterUserReq struct {
	Email              string `json:"email"`
	Password           string `json:"password"`
	UserName           string `json:"username"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Phone              string `json:"phone"`
	OrganizationDomain string `json:"organizationDomain"`
}

type RegisterUserRes struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	UserName     string `json:"username"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Phone        string `json:"phone"`
	Organization string `json:"organization"`
}
