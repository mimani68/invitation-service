package service

import (
	"context"

	"pluseid.io/invitation/data/entity/model"
)

type Invitation interface {
	GetAllInvitations(ctx context.Context) ([]model.Invitation, error)
}
