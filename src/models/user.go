package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	USER_ID          string = "id"
	USER_USERNAME    string = "username"
	USER_EMAIL       string = "email"
	USER_PASSWORD    string = "password"
	USER_FIRST_NAME  string = "first_name"
	USER_LAST_NAME   string = "last_name"
	USER_WRONG_PASS  string = "wrong_pass"
	USER_PERMISSIONS string = "permissions"
	USER_CREATED_AT  string = "created_at"
	USER_UPDATED_AT  string = "updated_at"
	USER_DELETED_AT  string = "deleted_at"
)

type User struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Username    string         `gorm:"uniqueIndex;not null" json:"username"`
	Email       string         `gorm:"not null" json:"email"`
	Password    string         `json:"-"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	WrongPass   int            `json:"-"`
	Permissions []string       `json:"permissions"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (model *User) TableName() string {
	return "users"
}
