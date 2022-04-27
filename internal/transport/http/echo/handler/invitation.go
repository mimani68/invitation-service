package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"pluseid.io/invitation/data/entity/model"
)

func (h *Handler) GetListOfInvitations(ctx echo.Context) error {
	list, err := h.InvitationSrv.GetAllInvitations(ctx.Request().Context())
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, list)
}

func (h *Handler) GetSingleInvitationById(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"msg": "your id is invalid"})
	}
	item, err := h.InvitationSrv.GetInvitationById(ctx.Request().Context(), int(id))
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, item)
}

func (h *Handler) CreateNewInvitation(ctx echo.Context) error {
	var newItem model.Invitation
	if err := ctx.Bind(&newItem); err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"msg": "request payload is not compatible"})
	}
	item, err := h.InvitationSrv.CreateInvitation(ctx.Request().Context(), newItem)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, item)
}

func (h *Handler) ExtendCodeExpireation(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"msg": "your id is invalid"})
	}
	item, err := h.InvitationSrv.ExtendInvitationById(ctx.Request().Context(), int(id))
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, item)
}

func (h *Handler) VerifyInvitation(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"msg": "your id is invalid"})
	}
	isValid := h.InvitationSrv.VerifyInvitationByCode(ctx.Request().Context(), int(id))
	return ctx.JSON(http.StatusOK, isValid)
}

func (h *Handler) UpdateInvitation(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"msg": "your id is invalid"})
	}
	var updatedCode model.Invitation
	if err := ctx.Bind(&updatedCode); err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"msg": "request payload is not compatible"})
	}
	item, err := h.InvitationSrv.UpdateInvitationById(ctx.Request().Context(), int(id), updatedCode)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, item)
}

func (h *Handler) InactiveInvitation(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    "125",
		"isValid": true,
	})
}

func (h *Handler) DeleteInvitation(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"msg": "your id is invalid"})
	}
	item, err := h.InvitationSrv.DeleteInvitationById(ctx.Request().Context(), int(id))
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, item)
}
