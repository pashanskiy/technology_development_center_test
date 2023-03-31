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

func (receiver *TransactionsService) CreateDeposit(ctx context.Context, request *apiTransactionService.CreateDepositRequest) (*apiCore.Uid, error) {
	logger := zerolog.Ctx(ctx)

	userUID, err := converter.String2UUID(logger, &request.UserUid.Uid)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert string to uuid")

		return nil, err
	}

	decimalMoney, err := decimal.NewFromString(request.Money.Money)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert money to decimal")

		return nil, err
	}
	transactionEntity := &entity.UserDeposit{
		ToUserUID: userUID,
		Money:     decimalMoney,
		Status:    entity.TransactionStatusCreated,
	}

	if err := receiver.db.Create(transactionEntity).Error; err != nil {
		logger.Error().Err(err).Msgf("failed create deposit transaction")

		return nil, errService.ErrCreateDeposit
	}

	logger.Trace().Msgf("deposit transaction successfully created. uid: %s", transactionEntity.UID.String())

	return &apiCore.Uid{Uid: transactionEntity.UID.String()}, nil
}
