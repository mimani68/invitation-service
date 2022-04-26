package system

import (
	"context"

	"pluseid.io/invitation/config"
	"pluseid.io/invitation/data/repository"
	"pluseid.io/invitation/data/service"
	"pluseid.io/invitation/internal/pkg/logger"
)

type system struct {
	cfg    config.Account
	mysql  repository.Mysql
	logger logger.Logger
}

func New(cfg config.Account, mysql repository.Mysql, logger logger.Logger) service.System {
	return &system{
		cfg:    cfg,
		mysql:  mysql,
		logger: logger,
	}
}

func (a *system) Ping(ctx context.Context) (string, error) {
	return "PONG", nil
}
