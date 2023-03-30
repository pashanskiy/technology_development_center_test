package transactions

import (
	"context"
	apiCore "technology_development_center_test/api/core"
	"technology_development_center_test/internal/db/entity"
	errService "technology_development_center_test/internal/transactions-service/errors"
	"technology_development_center_test/pkg/converters"

	"github.com/rs/zerolog"
)

func (receiver *TransactionsService) Delete(ctx context.Context, request *apiCore.Uid) (*apiCore.Empty, error) {
	logger := zerolog.Ctx(ctx)

	transactionUID, err := converters.String2UUID(logger, &request.Uid)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert string to uuid")

		return nil, err
	}

	transactionEntity := &entity.UserTransaction{
		UID: transactionUID,
	}

	if dbErr := receiver.db.Delete(transactionEntity, transactionEntity).Error; dbErr != nil {
		logger.Error().Err(dbErr).Msg("failed delete transaction")

		return nil, errService.ErrDeleteTransaction
	}

	logger.Trace().Msgf("transaction successfully deleted. transactionUID: %s", transactionUID.String())

	return &apiCore.Empty{}, nil
}
