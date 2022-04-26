package account

import (
	"context"

	"pluseid.io/invitation/config"
	"pluseid.io/invitation/data/entity/model"
	"pluseid.io/invitation/data/repository"
	"pluseid.io/invitation/data/service"
	"pluseid.io/invitation/internal/pkg/logger"
)

type acc struct {
	cfg    config.Account
	mysql  repository.Mysql
	logger logger.Logger
}

func New(cfg config.Account, mysql repository.Mysql, logger logger.Logger) service.Account {
	return &acc{
		cfg:    cfg,
		mysql:  mysql,
		logger: logger,
	}
}

func (a *acc) Login(ctx context.Context, username, password string) (string, error) {
	// start span with context

	user, err := a.mysql.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	_ = user
	// check password

	// create token
	token := ""

	return token, nil
}

func (a *acc) Register(ctx context.Context, user *model.User) (string, error) {
	return "", nil
}
