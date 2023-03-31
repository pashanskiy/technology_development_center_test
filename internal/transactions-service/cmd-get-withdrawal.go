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

func (receiver *TransactionsService) GetWithdrawal(ctx context.Context, request *apiTransactionService.TransactionInfo) (*apiTransactionService.TransactionsInfo, error) {
	logger := zerolog.Ctx(ctx)

	transactionUID, err := converters.String2UUID(logger, request.Uid)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert string to uuid")

		return nil, err
	}

	fromUserUID, err := converters.String2UUID(logger, request.FromUserUid)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert string to uuid")

		return nil, err
	}

	toUserUID, err := converters.String2UUID(logger, request.ToUserUid)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert string to uuid")

		return nil, err
	}

	withdrawalEntity := &entity.UserWithdrawal{
		UID:         transactionUID,
		FromUserUID: fromUserUID,
		ToUserUID:   toUserUID,
		Money:       decimal.Decimal{},
		CreatedAt:   time.Time{},
	}

	withdrawalEntities := &[]entity.UserWithdrawal{}

	if dbErr := receiver.db.Find(withdrawalEntities, withdrawalEntity).Error; dbErr != nil {
		logger.Error().Err(dbErr).Msg("failed get transaction")

		return nil, errService.ErrGetTransaction
	}

	transactionsList := make([]*apiTransactionService.TransactionInfo, len(*withdrawalEntities))

	for idx, transaction := range *withdrawalEntities {
		stringUID := transaction.UID.String()
		date := transaction.CreatedAt.UnixNano()
		transactionsList[idx] = &apiTransactionService.TransactionInfo{
			Uid:       &stringUID,
			ToUserUid: request.ToUserUid,
			Money:     &core.Money{Money: transaction.Money.String()},
			Status:    &transaction.Status,
			Created:   &date,
		}
	}

	logger.Trace().Msgf("transaction successfully was taken. transactions len: %d", len(transactionsList))

	return &apiTransactionService.TransactionsInfo{
		TransactionsInfo: transactionsList,
	}, nil
}
