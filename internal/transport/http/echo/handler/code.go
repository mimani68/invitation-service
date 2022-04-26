package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetListOfCodes(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, []map[string]interface{}{
		{
			"code":    "125",
			"isValid": true,
		},
		{
			"code":    "125",
			"isValid": true,
		},
	})
}

func (h *Handler) GetSingleCodeById(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    "125",
		"isValid": true,
	})
}

func (h *Handler) CreateNewCode(ctx echo.Context) error {
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

func (h *Handler) VerifyCode(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    "125",
		"isValid": true,
	})
}

func (h *Handler) UpdateCode(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    "125",
		"isValid": true,
	})
}

func (h *Handler) InactiveCode(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    "125",
		"isValid": true,
	})
}

func (h *Handler) DeleteCode(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    "125",
		"isValid": true,
	})
}
