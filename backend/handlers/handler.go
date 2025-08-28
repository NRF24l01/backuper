package handlers

import (
	"github.com/nrf24l01/go-web-utils/s3util"
	"gorm.io/gorm"

	"github.com/nrf24l01/backuper/core"
)

type Handler struct {
	DB *gorm.DB
	MinIOClient *s3util.Client
	Config *core.Config
}
