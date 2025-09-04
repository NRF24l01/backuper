package models

import (
	"github.com/google/uuid"
	"github.com/nrf24l01/go-web-utils/goorm"
)

type WorkerCapability struct {
	goorm.BaseModel
	WorkerID       uuid.UUID  `gorm:"column:worker_id"`
	Worker         *Worker    `gorm:"foreignKey:WorkerID"`
	Type           string     `gorm:"type:varchar(100)"`
	About          string     `gorm:"type:jsonb"`
	ToBackup       bool       `gorm:"default:false"`
	BackupInterval uint64     `gorm:"default:0"`
	LastBck        uint64     `gorm:"default:0"`
}
