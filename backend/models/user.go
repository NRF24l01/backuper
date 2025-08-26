package models

import (
	"github.com/NRF24l01/go-web-utils/goorm"
)

type User struct {
	goorm.BaseModel
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}