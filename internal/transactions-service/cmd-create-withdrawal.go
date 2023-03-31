package transactions

import (
	"context"
	apiCore "technology_development_center_test/api/core"
	apiTransactionService "technology_development_center_test/api/transactions"
	entity "technology_development_center_test/internal/db/entity"
	converter "technology_development_center_test/internal/transactions-service/converter"
	errService "technology_development_center_test/internal/transactions-service/errors"

	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

func (receiver *TransactionsService) CreateWithdrawal(ctx context.Context, request *apiTransactionService.CreateWithdrawalRequest) (*apiCore.Uid, error) {
	logger := zerolog.Ctx(ctx)

	fromUserUID, err := converter.String2UUID(logger, &request.FromUserUid.Uid)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert string to uuid")

		return nil, err
	}

	toUserUID, err := converter.String2UUID(logger, &request.ToUserUid.Uid)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert string to uuid")

		return nil, err
	}

	decimalMoney, err := decimal.NewFromString(request.Money.Money)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert money to decimal")

		return nil, err
	}
	transactionEntity := &entity.UserWithdrawal{
		FromUserUID: fromUserUID,
		ToUserUID:   toUserUID,
		Money:       decimalMoney,
		Status:      entity.TransactionStatusCreated,
	}

	if err := receiver.db.Create(transactionEntity).Error; err != nil {
		logger.Error().Err(err).Msgf("failed create withdrawal transaction")

		return nil, errService.ErrCreateDeposit
	}

	logger.Trace().Msgf("transaction successfully created. uid: %s", transactionEntity.UID.String())

	return &apiCore.Uid{Uid: transactionEntity.UID.String()}, nil
}
