package account

import (
	"context"

	"pluseid.io/invitation/config"
	"pluseid.io/invitation/data/entity/model"
	"pluseid.io/invitation/data/repository"
	"pluseid.io/invitation/data/service"
	"pluseid.io/invitation/internal/pkg/logger"
)

type accountSrv struct {
	cfg    config.Account
	mysql  repository.Mysql
	logger logger.Logger
}

func New(cfg config.Account, mysql repository.Mysql, logger logger.Logger) service.Account {
	return &accountSrv{
		cfg:    cfg,
		mysql:  mysql,
		logger: logger,
	}
}

func (a *accountSrv) Login(ctx context.Context, username, password string) (string, error) {
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

func (a *accountSrv) Register(ctx context.Context, user *model.User) (string, error) {
	return "", nil
}
