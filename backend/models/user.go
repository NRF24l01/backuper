package models

import (
	"github.com/labstack/gommon/log"
	"github.com/nrf24l01/go-web-utils/goorm"
	"github.com/nrf24l01/go-web-utils/passhash"
)

type User struct {
	goorm.BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
}


func (u *User) CheckPassword(password string) bool {
	res, err := passhash.CheckPassword(password, u.Password)
	log.Printf("Password check result: %v, error: %v", res, err)
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