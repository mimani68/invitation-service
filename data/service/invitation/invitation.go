package invitation

import (
	"context"
	"time"

	"pluseid.io/invitation/config"
	"pluseid.io/invitation/data/entity/model"
	"pluseid.io/invitation/data/repository"
	"pluseid.io/invitation/data/service"
	"pluseid.io/invitation/internal/pkg/logger"
	"pluseid.io/invitation/pkg/keygen"
)

type invitationSrv struct {
	cfg    config.Account
	mysql  repository.Mysql
	logger logger.Logger
}

func New(cfg config.Account, mysql repository.Mysql, logger logger.Logger) service.Invitation {
	return &invitationSrv{
		cfg:    cfg,
		mysql:  mysql,
		logger: logger,
	}
}

func (i *invitationSrv) GetAllInvitations(ctx context.Context) ([]model.Invitation, error) {
	invitationList, err := i.mysql.GetAllInvitations(ctx)
	if err != nil {
		return []model.Invitation{}, err
	}
	return invitationList, nil
}

func (i *invitationSrv) GetInvitationById(ctx context.Context, id int) (model.Invitation, error) {
	invitation, err := i.mysql.GetInvitationById(ctx, id)
	if err != nil {
		return model.Invitation{}, err
	}
	return invitation, nil
}

func (i *invitationSrv) CreateInvitation(ctx context.Context, newItem model.Invitation) (model.Invitation, error) {
	newItem.ID = 0
	newItem.Code = keygen.CodeGenerator(10)
	newItem.Expire = time.Now().Add(7 * 24 * time.Hour).Format(time.RFC3339)
	invitation, err := i.mysql.CreateNewInviteCode(ctx, newItem)
	if err != nil {
		return model.Invitation{}, err
	}
	return invitation, nil
}

func (i *invitationSrv) UpdateInvitationById(ctx context.Context, id int, updatedItem model.Invitation) (model.Invitation, error) {
	updatedItem.Code = keygen.CodeGenerator(10)
	updatedItem.Expire = time.Now().Add(7 * 24 * time.Hour).Format(time.RFC3339)
	invitation, err := i.mysql.UpdateInvitation(ctx, id, updatedItem)
	if err != nil {
		return model.Invitation{}, err
	}
	return invitation, nil
}

func (i *invitationSrv) DeleteInvitationById(ctx context.Context, id int) (model.Invitation, error) {
	invitation, err := i.mysql.DeleteInvitation(ctx, id)
	if err != nil {
		return model.Invitation{}, err
	}
	return invitation, nil
}

func (i *invitationSrv) ExtendInvitationById(ctx context.Context, id int) (model.Invitation, error) {
	item, err := i.GetInvitationById(ctx, id)
	if err != nil {
		return model.Invitation{}, err
	}
	item.Code = keygen.CodeGenerator(10)
	updatedItem, errUpdate := i.UpdateInvitationById(ctx, id, item)
	if errUpdate != nil {
		return model.Invitation{}, errUpdate
	}
	return updatedItem, nil
}

func (i *invitationSrv) RecalldInvitationById(ctx context.Context, id int) (model.Invitation, error) {
	item, err := i.GetInvitationById(ctx, id)
	if err != nil {
		return model.Invitation{}, err
	}
	item.IsActive = false
	updatedItem, errUpdate := i.UpdateInvitationById(ctx, id, item)
	if errUpdate != nil {
		return model.Invitation{}, errUpdate
	}
	return updatedItem, nil
}

func (i *invitationSrv) VerifyInvitationByCode(ctx context.Context, id int) bool {
	_, err := i.GetInvitationById(ctx, id)
	return err != nil
}
