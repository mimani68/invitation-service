package repository

import (
	"context"

	"pluseid.io/invitation/data/entity/model"
)

type Mysql interface {
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error

	GetAllInvitations(ctx context.Context) ([]model.Invitation, error)
	GetInvitationById(ctx context.Context, invitationId int) (model.Invitation, error)
	CreateNewInviteCode(ctx context.Context, newItem model.Invitation) (model.Invitation, error)
	UpdateInvitation(ctx context.Context, invitationId int, updatedItem model.Invitation) (model.Invitation, error)
	DeleteInvitation(ctx context.Context, invitationId int) (model.Invitation, error)
}
