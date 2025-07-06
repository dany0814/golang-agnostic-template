package repository

import (
	"context"
	entity "golang-agnostic-template/src/application/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) (res *entity.User, err error)
}
