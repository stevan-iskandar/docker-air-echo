package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	PERMISSION_ID         string = "id"
	PERMISSION_CODE       string = "code"
	PERMISSION_CREATED_AT string = "created_at"
	PERMISSION_UPDATED_AT string = "updated_at"
	PERMISSION_DELETED_AT string = "deleted_at"
)

type Permission struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Code      string         `gorm:"uniqueIndex;not null" json:"code"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (model *Permission) TableName() string {
	return "permissions"
}
