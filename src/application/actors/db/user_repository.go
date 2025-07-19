package surreal

import (
	"context"
	"errors"
	"fmt"
	entity "golang-agnostic-template/src/application/domain/model"
	"golang-agnostic-template/src/application/domain/utils"
	"golang-agnostic-template/src/pkg/database"
	"strings"

	"github.com/surrealdb/surrealdb.go"
	"github.com/surrealdb/surrealdb.go/pkg/models"
)

type UserRepository struct {
	conn database.IDatabaseConnection
}

func NewUserRepository(conn database.IDatabaseConnection) *UserRepository {
	return &UserRepository{conn: conn}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) (res *entity.User, err error) {
	response, err := surrealdb.Create[entity.User, models.Table](r.conn.DB(), models.Table("users"), &user)
	return response, err
}

func (r *UserRepository) Read(ctx context.Context, email *string, username *string, id *string) (res *entity.User, err error) {
	var conditions []string
	var args = make(map[string]interface{})

	if email != nil {
		conditions = append(conditions, "email = $email")
		args["email"] = *email
	}
	if id != nil {
		conditions = append(conditions, "id = $id")
		args["id"] = &models.RecordID{
			ID:    *id,
			Table: "users",
		}
	}
	if username != nil {
		conditions = append(conditions, "username = $username")
		args["username"] = *username
	}

	if len(conditions) == 0 {
		return nil, errors.New("at least one parameter (email, id, or username) is required")
	}

	query := fmt.Sprintf("SELECT * FROM users WHERE %s", strings.Join(conditions, " AND "))
	response, err := surrealdb.Query[[]entity.User](r.conn.DB(), query, args)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if len(*response) == 0 || len((*response)[0].Result) == 0 {
		return nil, errors.New(utils.ErrUserNotFound)
	}

	userFounded := (*response)[0].Result[0]
	return &userFounded, nil
}

func (r *UserRepository) Update(ctx context.Context, id *string, user *entity.User) (res *entity.User, err error) {
	userId := models.RecordID{
		ID:    *id,
		Table: "users",
	}
	user.ID = &userId
	response, err := surrealdb.Update[entity.User, models.RecordID](r.conn.DB(), userId, &user)
	return response, err
}

func (r *UserRepository) Delete(ctx context.Context, id *string) (res *entity.User, err error) {
	userId := models.RecordID{
		ID:    *id,
		Table: "users",
	}
	response, err := surrealdb.Delete[entity.User, models.RecordID](r.conn.DB(), userId)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if response.ID == nil {
		return nil, fmt.Errorf("User not founded to delete: %s", *id)
	}
	return response, err
}
