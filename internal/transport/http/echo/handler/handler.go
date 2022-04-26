package handler

import (
	"pluseid.io/invitation/data/service"
	"pluseid.io/invitation/internal/pkg/logger"
)

type Handler struct {
	Logger  logger.Logger
	Account service.Account
	System  service.System
}
