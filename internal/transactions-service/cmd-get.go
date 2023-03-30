package transactions

import (
	"context"
	"technology_development_center_test/api/core"
	apiTransactionService "technology_development_center_test/api/transactions"
	"technology_development_center_test/internal/db/entity"
	errService "technology_development_center_test/internal/transactions-service/errors"
	"technology_development_center_test/pkg/converters"
	"time"

	"github.com/shopspring/decimal"

	"github.com/rs/zerolog"
)

func (receiver *TransactionsService) Get(ctx context.Context, request *apiTransactionService.TransactionInfo) (*apiTransactionService.TransactionsInfo, error) {
	logger := zerolog.Ctx(ctx)

	transactionUID, err := converters.String2UUID(logger, request.Uid)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert string to uuid")

		return nil, err
	}

	transactionEntity := &entity.UserTransaction{
		UID:       transactionUID,
		Money:     decimal.Decimal{},
		CreatedAt: time.Time{},
	}

	transactionEntities := &[]entity.UserTransaction{}

	if dbErr := receiver.db.Find(transactionEntities, transactionEntity).Error; dbErr != nil {
		logger.Error().Err(dbErr).Msg("failed get transaction")

		return nil, errService.ErrGetTransaction
	}

	transactionsList := make([]*apiTransactionService.TransactionInfo, len(*transactionEntities))

	for idx, transaction := range *transactionEntities {
		stringUID := transaction.UID.String()
		date := transaction.CreatedAt.UnixNano()
		transactionsList[idx] = &apiTransactionService.TransactionInfo{
			Uid:     &stringUID,
			Money:   &core.Money{Money: transaction.Money.String()},
			Created: &date,
		}
	}

	logger.Trace().Msgf("transaction successfully was taken. transactions len: %d", len(transactionsList))

	return &apiTransactionService.TransactionsInfo{
		TransactionsInfo: transactionsList,
	}, nil
}
