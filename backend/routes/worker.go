package routes

import (
	"github.com/nrf24l01/backuper/handlers"
	"github.com/nrf24l01/backuper/schemas"
	"github.com/nrf24l01/go-web-utils/echokit"

	"github.com/labstack/echo/v4"
)

func RegisterWorkerRoutes(e *echo.Echo, h *handlers.Handler) {
	group := e.Group("/workers")
	group.Use(echokit.JWTMiddleware([]byte(h.Config.JWTAccessSecret)))

	group.POST("", h.WorkerCreateHandler, echokit.ValidationMiddleware(func() interface{} {
			return &schemas.WorkerCreateRequest{}
		}))
}