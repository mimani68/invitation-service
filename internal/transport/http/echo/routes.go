package echo

import (
	"github.com/labstack/echo/v4/middleware"
	"pluseid.io/invitation/pkg/auth"
)

func (r *rest) routing() {

	r.echo.GET("/ping", r.handler.Ping)

	console := r.echo.Group("console/v1", middleware.BasicAuth(auth.BasicAuthentication))
	// console.POST("/v1/login", r.handler.Login)
	// console.POST("/v1/register", r.handler.Register)

	codeRoute := console.Group("/codes")
	codeRoute.GET("", r.handler.GetListOfInvitations)
	codeRoute.GET("/:codeId", r.handler.GetSingleInvitationById)
	codeRoute.POST("", r.handler.CreateNewInvitation)
	codeRoute.PUT("/:codeId", r.handler.UpdateInvitation)
	codeRoute.POST("/:codeId/inactive", r.handler.InactiveInvitation)
	codeRoute.DELETE("/:codeId", r.handler.DeleteInvitation)

	public := r.echo.Group("v1")
	public.POST("/code", r.handler.VerifyInvitation)

}
