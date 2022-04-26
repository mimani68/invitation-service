package echo

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pluseid.io/invitation/data/service"
	"pluseid.io/invitation/internal/pkg/logger"
	"pluseid.io/invitation/internal/transport/http"
	"pluseid.io/invitation/internal/transport/http/echo/handler"
)

type rest struct {
	echo    *echo.Echo
	handler *handler.Handler
}

func New(logger logger.Logger, accSrv service.Account, systemSrv service.System, invitationSrv service.Invitation) http.Rest {
	return &rest{
		echo: echo.New(),
		handler: &handler.Handler{
			Logger:        logger,
			AccountSrv:    accSrv,
			SystemSrv:     systemSrv,
			InvitationSrv: invitationSrv,
		}}
}

func (r *rest) Start(address string) error {
	r.echo.Use(middleware.Recover())

	r.routing()

	return r.echo.Start(address)
}

func (r *rest) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // use config for time
	defer cancel()

	return r.echo.Shutdown(ctx)
}
