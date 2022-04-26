package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"pluseid.io/invitation/internal/transport/http/response"
)

func (h *Handler) Ping(c echo.Context) error {
	pingString, _ := h.SystemSrv.Ping(c.Request().Context())
	return c.JSON(http.StatusOK, response.Ping{Message: pingString})
}
