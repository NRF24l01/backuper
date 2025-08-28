package models

import (
	"github.com/nrf24l01/go-web-utils/goorm"
	"github.com/nrf24l01/go-web-utils/passhash"
)

type User struct {
	goorm.BaseModel
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}


func (u *User) CheckPassword(password string) bool {
	res, err := passhash.CheckPassword(u.Password, password)
	return res && err == nil
}

func (u *User) HashPassword() error {
	hash, err := passhash.HashPassword(u.Password, passhash.DefaultParams)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}