package converter

import (
	// apiTransactionService "technology_development_center_test/api/transactions"

	"github.com/gofrs/uuid/v5"
	"github.com/rs/zerolog"
)

// func TransactionStatus2String(logger *zerolog.Logger, status apiTransactionService.TransactionStatus) string {
// 	logger.Debug().Msgf("status: %s", status)

// 	if apiTransactionService.TransactionStatus_TRANSACTION_STATUS_UNSPECIFIED == status {
// 		return ""
// 	}

// 	result := status.String()[19:]

// 	logger.Debug().Msgf("result: %s", result)

// 	return result
// }

// func String2TransactionsStatus(logger *zerolog.Logger, value string) apiTransactionService.TransactionStatus {
// 	logger.Debug().Msgf("value: %s", value)

// 	code := strings.TrimSpace(value)
// 	if code == "" {
// 		return apiTransactionService.TransactionStatus_TRANSACTION_STATUS_UNSPECIFIED
// 	}

// 	enumValue := fmt.Sprintf("STATUS_%s", strings.ToUpper(code))
// 	logger.Debug().Msgf("enumValue: %s", enumValue)

// 	result, exist := apiTransactionService.TransactionStatus_value[enumValue]
// 	if !exist {
// 		logger.Error().Msgf("Status is not found: %s", value)

// 		return apiTransactionService.TransactionStatus_TRANSACTION_STATUS_UNSPECIFIED
// 	}

// 	return *apiTransactionService.TransactionStatus(result).Enum()
// }

func String2UUID(logger *zerolog.Logger, stringUUID *string) (*uuid.UUID, error) {
	if stringUUID == nil {

		return nil, nil
	}

	rawuuid, err := uuid.FromString(*stringUUID)
	if err != nil {
		logger.Error().Err(err).Msg("failed parse userUID")

		return nil, ErrInvalidUID
	}

	return &rawuuid, nil
}
