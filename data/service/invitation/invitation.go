package account

import (
	"context"

	"pluseid.io/invitation/config"
	"pluseid.io/invitation/data/entity/model"
	"pluseid.io/invitation/data/repository"
	"pluseid.io/invitation/data/service"
	"pluseid.io/invitation/internal/pkg/logger"
)

type InvitationService struct {
	cfg    config.Account
	mysql  repository.Mysql
	logger logger.Logger
}

func New(cfg config.Account, mysql repository.Mysql, logger logger.Logger) service.Invitation {
	return &InvitationService{
		cfg:    cfg,
		mysql:  mysql,
		logger: logger,
	}
}

func (a *InvitationService) GetAllInvitations(ctx context.Context) ([]model.Invitation, error) {
	invitationList, err := a.mysql.GetAllInvitations(ctx)
	if err != nil {
		return []model.Invitation{}, err
	}
	return invitationList, nil
}
