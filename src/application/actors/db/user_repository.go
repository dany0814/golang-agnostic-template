package surreal

import (
	"context"
	"fmt"
	entity "golang-agnostic-template/src/application/domain/model"
	"golang-agnostic-template/src/pkg/database"

	"github.com/surrealdb/surrealdb.go"
	"github.com/surrealdb/surrealdb.go/pkg/models"
)

type UserRepositorySurreal struct {
	conn database.IDatabaseConnection
}

func NewUserRepositorySurreal(conn database.IDatabaseConnection) *UserRepositorySurreal {
	return &UserRepositorySurreal{conn: conn}
}

func (r *UserRepositorySurreal) Create(ctx context.Context, user entity.User) (res *entity.User, err error) {
	response, err := surrealdb.Create[entity.User](r.conn.DB(), models.Table("users"), user)
	if err != nil {
		fmt.Printf("failed to create user: %w", err)
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	if response == nil {
		fmt.Printf("surrealdb.Create returned nil response: %w", response)
		return nil, fmt.Errorf("surrealdb.Create returned nil response")
	}

	fmt.Printf("Created user: %+v\n", response)
	return response, nil
}
