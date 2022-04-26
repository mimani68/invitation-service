package account

import (
	"context"

	"pluseid.io/invitation/config"
	"pluseid.io/invitation/data/repository"
	"pluseid.io/invitation/data/service"
	"pluseid.io/invitation/internal/pkg/logger"
)

type codeService struct {
	cfg    config.Account
	mysql  repository.Mysql
	logger logger.Logger
}

func New(cfg config.Account, mysql repository.Mysql, logger logger.Logger) service.Code {
	return &codeService{
		cfg:    cfg,
		mysql:  mysql,
		logger: logger,
	}
}

func (a *codeService) GetAllCodes(ctx context.Context) (string, error) {
	user, err := a.mysql.GetUserByUsername(ctx, "as")
	if err != nil {
		return "", err
	}
	_ = user
	// check password

	// create token
	token := ""

	return token, nil
}
