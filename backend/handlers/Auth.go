package handlers

import (
	"net/http"

	"github.com/NRF24l01/backuper/models"
	"github.com/NRF24l01/backuper/schemas"
	"github.com/labstack/echo/v4"
)

func (h* Handler) UserLoginHandler(c echo.Context) error {
	user_data := c.Get("validatedBody").(*schemas.AuthRequest)

	var user models.User
	if err := h.DB.Where("username = ?", user_data.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, schemas.Message{Status: "Invalid username or password"})
	}

	if !user.CheckPassword(user_data.Password) {
		return c.JSON(http.StatusUnauthorized, schemas.Message{Status: "Invalid username or password"})
	}

	return c.JSON(http.StatusOK, schemas.Message{Status: "User logged in successfully"})
}