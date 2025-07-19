package dto

type GetUserRes struct {
	ID                   string `json:"id"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	Username             string `json:"username"`
	Phone                string `json:"phone"`
	State                string `json:"state"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
	DeletedAt            string `json:"deleted_at"`
	NotificationsByEmail bool   `json:"notification_by_email"`
	NotificationsBySms   bool   `json:"notification_by_sms"`
}
