package entity

import (
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/shopspring/decimal"

	"gorm.io/gorm"
)

type UserWithdrawals struct {
	UID *uuid.UUID `gorm:"column:uid; type:uuid; primaryKey"`

	UserUID string          `gorm:"column:user_uid"`
	Money   decimal.Decimal `gorm:"column:money"`
	Status  string          `gorm:"column:status"`

	CreatedAt time.Time      `gorm:"column:created_at; <-:create"`
	UpdatedAt time.Time      `gorm:"column:updated_at; autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (UserWithdrawals) TableName() string { return "user_withdrawals" }
