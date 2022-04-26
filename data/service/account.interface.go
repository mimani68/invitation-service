package service

import (
	"context"

	"pluseid.io/invitation/data/entity/model"
)

type Account interface {
	Login(ctx context.Context, username, password string) (string, error)
	Register(ctx context.Context, user *model.User) (string, error)
}
