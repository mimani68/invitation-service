package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"pluseid.io/invitation/internal/transport/http/request"
	"pluseid.io/invitation/internal/transport/http/response"
)

func (h *Handler) Login(c echo.Context) error {
	// start span with context

	req := new(request.Login)

	if err := c.Bind(req); err != nil {
		h.Logger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest})
	}

	token, err := h.AccountHandler.Login(c.Request().Context(), req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest}) // i know can better error handling
	}

	return c.JSON(http.StatusOK, response.Login{Token: token, Status: http.StatusOK})
}

func (h *Handler) Register(c echo.Context) error {
	return nil
}
