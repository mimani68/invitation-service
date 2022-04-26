package echo

import (
	"github.com/labstack/echo/v4/middleware"
	"pluseid.io/invitation/pkg/auth"
)

func (r *rest) routing() {

	console := r.echo.Group("console/v1", middleware.BasicAuth(auth.BasicAuthentication))
	console.GET("/ping", r.handler.Ping)
	// console.POST("/v1/login", r.handler.Login)
	// console.POST("/v1/register", r.handler.Register)

	codeRoute := console.Group("/codes")
	codeRoute.GET("", r.handler.GetListOfCodes)
	codeRoute.GET("/:codeId", r.handler.GetSingleCodeById)
	codeRoute.POST("", r.handler.CreateNewCode)
	codeRoute.PUT("/:codeId", r.handler.UpdateCode)
	codeRoute.POST("/:codeId/inactive", r.handler.InactiveCode)
	codeRoute.DELETE("/:codeId", r.handler.DeleteCode)

	public := r.echo.Group("v1")
	public.POST("/code", r.handler.VerifyCode)

}
