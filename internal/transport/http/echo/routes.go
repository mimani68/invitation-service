package echo

import (
	"crypto/subtle"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pluseid.io/invitation/config"
)

func (r *rest) routing() {
	console := r.echo.Group("console/v1", middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte(config.GetAdmin().Username)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(config.GetAdmin().Password)) == 1 {
			return true, nil
		}
		return false, nil
	}))
	// console.POST("/v1/login", r.handler.Login)
	// console.POST("/v1/register", r.handler.Register)
	// console.GET("/v1/ping", r.handler.Ping)
	console.GET("/ping", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "PONG")
	})
	console.GET("/codes", func(ctx echo.Context) error {
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
	})
	console.GET("/codes/:codeId", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    "125",
			"isValid": true,
		})
	})
	console.POST("/codes", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    "125",
			"isValid": true,
		})
	})
	console.PUT("/codes/:codeId", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    "125",
			"isValid": true,
		})
	})
	console.POST("/codes/:codeId/inactive", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    "125",
			"isValid": true,
		})
	})
	console.DELETE("/codes/:codeId", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    "125",
			"isValid": true,
		})
	})

	public := r.echo.Group("v1")
	public.GET("/ping", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "PONG")
	})
	public.POST("/code", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    "125",
			"isValid": true,
		})
	})

}
