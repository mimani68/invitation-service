package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetListOfInvitations(ctx echo.Context) error {
	list, _ := h.InvitationHandler.GetAllInvitations(ctx.Request().Context())
	return ctx.JSON(http.StatusOK, list)
}

func (h *Handler) GetSingleInvitationById(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    "125",
		"isValid": true,
	})
}

func (h *Handler) CreateNewInvitation(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    "125",
		"isValid": true,
	})
}

func (h *Handler) ExtendCodeExpireation(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    "125",
		"isValid": true,
	})
}

func (h *Handler) VerifyInvitation(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    "125",
		"isValid": true,
	})
}

func (h *Handler) UpdateInvitation(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    "125",
		"isValid": true,
	})
}

func (h *Handler) InactiveInvitation(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    "125",
		"isValid": true,
	})
}

func (h *Handler) DeleteInvitation(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    "125",
		"isValid": true,
	})
}
