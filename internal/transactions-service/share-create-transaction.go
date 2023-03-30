package transactions

import (
	apiCore "technology_development_center_test/api/core"
	apiTransactionService "technology_development_center_test/api/transactions"
	"technology_development_center_test/internal/db/entity"
	converter "technology_development_center_test/internal/transactions-service/converter"
	errService "technology_development_center_test/internal/transactions-service/errors"

	"github.com/shopspring/decimal"

	"github.com/rs/zerolog"
)

func (receiver *TransactionsService) createTransaction(logger *zerolog.Logger, request *apiTransactionService.CreateTransaction, transactionType TransactionType) (*apiCore.Uid, error) {

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
	transactionEntity := &entity.UserTransaction{
		UserUID: userUID,
		Money:   decimalMoney,
		Type:    string(transactionType),
		Status:  converter.TransactionStatus2String(logger, apiTransactionService.TransactionStatus_TRANSACTION_STATUS_CREATED),
	}

	if err := receiver.db.Create(transactionEntity).Error; err != nil {
		logger.Error().Err(err).Msgf("failed create %s transaction", string(transactionType))

		return nil, errService.ErrCreateDeposit
	}

	logger.Trace().Msgf("%s transaction successfully created. uid: %s", string(transactionType), transactionEntity.UID.String())

	return &apiCore.Uid{Uid: transactionEntity.UID.String()}, nil
}
