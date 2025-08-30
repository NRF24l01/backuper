package models

import (
	"crypto/rand"

	"github.com/nrf24l01/go-web-utils/goorm"
)

type Worker struct {
	goorm.BaseModel
	Name       string `json:"name"`
	Token      string `json:"token"`
	LastOnline int64  `json:"last_online" gorm:"default:0"`
}

func (w *Worker) NewToken() error {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return err
	}
	for i := range b {
		b[i] = letters[int(b[i])%len(letters)]
	}
	w.Token = string(b)
	return nil
}