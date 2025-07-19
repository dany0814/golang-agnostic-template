package repository

import (
	"context"
	entity "golang-agnostic-template/src/application/domain/model"
)

type IUserRepository interface {
	Create(ctx context.Context, user *entity.User) (res *entity.User, err error)
	Read(ctx context.Context, email *string, username *string, id *string) (res *entity.User, err error)
	Update(ctx context.Context, id *string, user *entity.User) (res *entity.User, err error)
	Delete(ctx context.Context, id *string) (res *entity.User, err error)
}
