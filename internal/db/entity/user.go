package entity

import (
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/shopspring/decimal"

	"gorm.io/gorm"
)

type UserEntity struct {
	UID *uuid.UUID `gorm:"column:uid; type:uuid; primaryKey; default:uuid_generate_v4()"`

	Money decimal.Decimal `gorm:"column:money"`

	CreatedAt time.Time      `gorm:"column:created_at; <-:create"`
	UpdatedAt time.Time      `gorm:"column:updated_at; autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (UserEntity) TableName() string { return "users" }
