package entity

import (
	"golang-agnostic-template/src/application/domain/utils"
	"net/mail"
	"time"

	"github.com/surrealdb/surrealdb.go/pkg/models"
)

type User struct {
	ID        *models.RecordID `json:"id"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	Username  string           `json:"username"`
	Phone     string           `json:"phone"`
	State     string           `json:"state"`
	CreatedAt string           `json:"created_at"`
	UpdatedAt string           `json:"updated_at"`
	DeletedAt string           `json:"deleted_at"`
	Settings  UserSettings     `json:"settings"`
}

type UserSettings struct {
	EmailNotifications bool   `json:"email"`
	SmsNotifications   bool   `json:"sms"`
	Language           string `json:"language"`
}

func (u *User) ValidateEmail() error {
	if &u.Email != nil {
		_, err := mail.ParseAddress(u.Email)
		if err != nil {
			return utils.ErrEmailUser
		}
	}
	return nil
}

func (u *User) BuildUser() {
	u.ID = &models.RecordID{
		ID: utils.UID().New(),
	}
	u.State = utils.ACTIVE
	u.CreatedAt = time.Now().String()
	u.UpdatedAt = time.Now().String()
	u.DeletedAt = time.Now().String()
}
