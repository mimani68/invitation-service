package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"pluseid.io/invitation/data/entity/model"
	"pluseid.io/invitation/data/repository"
)

func (h *Handler) GetListOfInvitations(ctx echo.Context) error {
	meta := repository.QueryMeta{
		IsActive: ctx.QueryParam("is_active") == "true",
	}
	if len(ctx.QueryParam("is_active")) > 2 {
		meta.Query = true
	}
	list, err := h.InvitationSrv.GetAllInvitations(ctx.Request().Context(), meta)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, list)
}

func (h *Handler) GetSingleInvitationById(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("codeId"), 10, 0)
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
	id, err := strconv.ParseInt(ctx.Param("codeId"), 10, 0)
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
	var verifiyPayloadRequest map[string]interface{}
	if err := ctx.Bind(&verifiyPayloadRequest); err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"msg": "request payload is not compatible", "valid": false})
	}
	if fmt.Sprintf("%T", verifiyPayloadRequest["code"]) != "string" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"msg": "request payload has invalid code", "valid": false})
	}
	if len(verifiyPayloadRequest["code"].(string)) < 3 {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"msg": "your code is very short", "valid": false})
	}
	isValid := h.InvitationSrv.VerifyInvitationByCode(ctx.Request().Context(), verifiyPayloadRequest["code"].(string))
	return ctx.JSON(http.StatusOK, map[string]interface{}{"valid": isValid})
}

func (h *Handler) UpdateInvitation(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("codeId"), 10, 0)
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
	id, err := strconv.ParseInt(ctx.Param("codeId"), 10, 0)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"msg": "your id is invalid"})
	}
	item, err := h.InvitationSrv.RecalldInvitationById(ctx.Request().Context(), int(id))
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, item)
}

func (h *Handler) DeleteInvitation(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("codeId"), 10, 0)
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
