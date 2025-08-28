package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nrf24l01/backuper/models"
	"github.com/nrf24l01/backuper/schemas"
	"github.com/nrf24l01/go-web-utils/jwtutil"
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

	accessToken, refreshToken, err := jwtutil.GenerateTokenPair(user.ID.String(), user.Username, []byte(h.Config.JWTAccessSecret), []byte(h.Config.JWTRefreshSecret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, schemas.DefaultInternalErrorResponse)
	}

	cookie := new(http.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = refreshToken
	cookie.HttpOnly = true
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, schemas.AuthResponse{AccessToken: accessToken})
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

	accessToken, refreshToken, err := jwtutil.GenerateTokenPair(newUser.ID.String(), newUser.Username, []byte(h.Config.JWTAccessSecret), []byte(h.Config.JWTRefreshSecret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, schemas.DefaultInternalErrorResponse)
	}

	cookie := new(http.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = refreshToken
	cookie.HttpOnly = true
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.JSON(http.StatusCreated, schemas.AuthResponse{AccessToken: accessToken})
}
