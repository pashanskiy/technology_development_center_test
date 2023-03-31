package entity

import (
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/shopspring/decimal"

	"gorm.io/gorm"
)

type UserDeposit struct {
	UID *uuid.UUID `gorm:"column:uid; type:uuid; primaryKey; default:uuid_generate_v4()"`

	ToUserUID *uuid.UUID `gorm:"column:to_user_uid"`

	Money  decimal.Decimal   `gorm:"column:money"`
	Status TransactionStatus `gorm:"column:status"`

	CreatedAt time.Time      `gorm:"column:created_at; <-:create"`
	UpdatedAt time.Time      `gorm:"column:updated_at; autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (UserDeposit) TableName() string { return "user_deposit" }
