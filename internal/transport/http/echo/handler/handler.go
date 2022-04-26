package handler

import (
	"pluseid.io/invitation/data/service"
	"pluseid.io/invitation/internal/pkg/logger"
)

type Handler struct {
	Logger            logger.Logger
	AccountHandler    service.Account
	SystemHandler     service.System
	InvitationHandler service.Invitation
}
