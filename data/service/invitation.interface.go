package service

import (
	"context"

	"pluseid.io/invitation/data/entity/model"
)

type Invitation interface {
	GetAllInvitations(ctx context.Context) ([]model.Invitation, error)
	GetInvitationById(ctx context.Context, id int) (model.Invitation, error)
	CreateInvitation(ctx context.Context, newItem model.Invitation) (model.Invitation, error)
	UpdateInvitationById(ctx context.Context, id int, updatedItem model.Invitation) (model.Invitation, error)
	DeleteInvitationById(ctx context.Context, id int) (model.Invitation, error)

	ExtendInvitationById(ctx context.Context, id int) (model.Invitation, error)
	RecalldInvitationById(ctx context.Context, id int) (model.Invitation, error)
	VerifyInvitationByCode(ctx context.Context, id int) bool
}
