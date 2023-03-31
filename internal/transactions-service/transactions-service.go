package transactions

import (
	apiTransactionsService "technology_development_center_test/api/transactions"

	"gorm.io/gorm"
)

type TransactionsService struct {
	apiTransactionsService.TransactionsServiceServer

	db *gorm.DB
}
func NewTransactionsService(db *gorm.DB) *TransactionsService {
	return &TransactionsService{db: db}
}


