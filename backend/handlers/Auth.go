package handlers

import (
	"net/http"

	"github.com/NRF24l01/backuper/schemas"
	"github.com/labstack/echo/v4"
)

func (h* Handler) UserLoginHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, schemas.Message{Status: "User logged in successfully"})
}