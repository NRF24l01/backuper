package routes

import (
	"github.com/nrf24l01/backuper/handlers"
	"github.com/nrf24l01/backuper/schemas"
	"github.com/nrf24l01/go-web-utils/echokit"

	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(e *echo.Echo, h *handlers.Handler) {
	group := e.Group("/auth")

	group.POST("/login", h.UserLoginHandler, echokit.ValidationMiddleware(func() interface{} {
		return &schemas.AuthRequest{}
	}))

	if (h.Config.TestEnv) {
		group.POST("/register", h.UserRegisterHandler, echokit.ValidationMiddleware(func() interface{} {
			return &schemas.AuthRequest{}
		}))
	}
}