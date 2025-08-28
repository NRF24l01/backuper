package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nrf24l01/backuper/models"
	"github.com/nrf24l01/backuper/schemas"
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

	

	// Set the refresh token in an HttpOnly cookie
	cookie := new(http.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = refreshToken
	cookie.HttpOnly = true
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, schemas.Message{Status: "User logged in successfully"})
}

func (h* Handler) UserRegisterHandler(c echo.Context) error {
	user_data := c.Get("validatedBody").(*schemas.AuthRequest)

	var existingUser models.User
	if err := h.DB.Where("username = ?", user_data.Username).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusConflict, schemas.Message{Status: "User already exists"})
	}

	newUser := models.User{
		Username: user_data.Username,
		Password: user_data.Password,
	}

	if err := newUser.HashPassword(); err != nil {
		return c.JSON(http.StatusInternalServerError, schemas.DefaultInternalErrorResponse)
	}

	if err := h.DB.Create(&newUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, schemas.DefaultInternalErrorResponse)
	}

	return c.JSON(http.StatusCreated, schemas.Message{Status: "User registered successfully"})
}
