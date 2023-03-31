package scheduler

import (
	"errors"
	"technology_development_center_test/internal/db/entity"

	"github.com/gofrs/uuid/v5"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func (receiver *TransactionSchedulerService) withdrawalJob() {
	logger := receiver.logger.With().Str("scheduler", "withdrawal").Logger()

	transactionUIDList, err := receiver.getWithdrawalUIDsToProcess(logger)
	if err != nil {
		logger.Error().Err(err).Msgf("failed get entities from database")
	}

	for _, transactionUID := range transactionUIDList {
		if err := receiver.jobProcessWithdrawal(logger, transactionUID); err != nil {
			logger.Error().Err(err).Msgf("failed withdrawal transaction; %s", transactionUID.String())
		}
	}

	logger.Info().Msgf("withdrawal job completed")
}

func (receiver *TransactionSchedulerService) jobProcessWithdrawal(logger zerolog.Logger, transactionUID uuid.UUID) error {
	tx := receiver.db.Begin()

	transactionEntity := entity.UserWithdrawal{UID: &transactionUID}
	if err := tx.Take(&transactionEntity).Error; err != nil {
		logger.Error().Err(err).Msgf("failed get user transaction entity from database")

		return err
	}

	fromUserEntity := entity.UserEntity{UID: transactionEntity.FromUserUID}
	if err := tx.Take(&fromUserEntity).Error; err != nil {
		logger.Error().Err(err).Msgf("failed get user entity from database")

		return err
	}

	toUserEntity := entity.UserEntity{UID: transactionEntity.ToUserUID}
	if err := tx.Take(&toUserEntity).Error; err != nil {
		logger.Error().Err(err).Msgf("failed get user entity from database")

		return err
	}

	if fromUserEntity.Money.LessThan(transactionEntity.Money) {
		logger.Info().Msgf("user %s not enough money to withdrawal", fromUserEntity.UID.String())

		transactionEntity.Status = entity.TransactionStatusRejected
		tx.Save(&transactionEntity)

	} else {

		fromUserEntity.Money = fromUserEntity.Money.Sub(transactionEntity.Money)

		if err := tx.Save(&fromUserEntity).Error; err != nil {
			logger.Error().Err(err).Msgf("failed save from user entity to database")

			return err
		}

		toUserEntity.Money = toUserEntity.Money.Add(transactionEntity.Money)

		if err := tx.Save(&toUserEntity).Error; err != nil {
			logger.Error().Err(err).Msgf("failed save to user entity to database")

			return err
		}

		transactionEntity.Status = entity.TransactionStatusSuccess
		if err := tx.Save(&transactionEntity).Error; err != nil {
			logger.Error().Err(err).Msgf("failed save user entity to database")

			return err
		}
	}
	tx.Commit()

	return nil
}

func (receiver *TransactionSchedulerService) getWithdrawalUIDsToProcess(logger zerolog.Logger) (confirmedEntitiesUIDs []uuid.UUID, err error) {
	if err := receiver.db.Model(&entity.UserWithdrawal{}).
		Where(&entity.UserDeposit{Status: entity.TransactionStatusCreated}).
		Order("created_at ASC").
		Select("uid").
		Find(&confirmedEntitiesUIDs).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Trace().Msg("entities not found")

		return nil, nil
	} else if err != nil {
		logger.Error().Err(err).Msg("failed to get uid in processing")

		return nil, err
	}

	return confirmedEntitiesUIDs, nil
}
