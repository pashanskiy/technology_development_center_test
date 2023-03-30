package transactions

import (
	"context"
	apiCore "technology_development_center_test/api/core"
	apiTransactionService "technology_development_center_test/api/transactions"

	"github.com/rs/zerolog"
)

func (receiver *TransactionsService) CreateDeposit(ctx context.Context, request *apiTransactionService.CreateTransaction) (*apiCore.Uid, error) {
	logger := zerolog.Ctx(ctx)

	uid, err := receiver.createTransaction(logger, request, Deposit)

	return &apiCore.Uid{Uid: uid.String()}, err
}
