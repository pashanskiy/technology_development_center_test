package scheduler

import (
	"errors"
	"technology_development_center_test/internal/db/entity"

	"github.com/gofrs/uuid/v5"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func (receiver *TransactionSchedulerService) depositJob() {
	logger := receiver.logger.With().Str("scheduler", "deposit").Logger()

	transactionUIDList, err := receiver.geDepositUIDsToProcess(logger)
	if err != nil {
		logger.Error().Err(err).Msgf("failed get entities from database")
	}

	for _, transactionUID := range transactionUIDList {
		if err := receiver.jobProcessDeposit(logger, transactionUID); err != nil {
			logger.Error().Err(err).Msgf("failed deposit transaction; %s", transactionUID.String())
		}
	}

	logger.Info().Msgf("deposit job completed")
}

func (receiver *TransactionSchedulerService) jobProcessDeposit(logger zerolog.Logger, transactionUID uuid.UUID) error {
	tx := receiver.db.Begin()

	depositEntity := entity.UserDeposit{UID: &transactionUID}
	if err := tx.Take(&depositEntity).Error; err != nil {
		logger.Error().Err(err).Msgf("failed get user transaction entity from database")

		return err
	}

	toUserEntity := entity.UserEntity{UID: depositEntity.ToUserUID}
	if err := tx.Take(&toUserEntity).Error; err != nil {
		logger.Error().Err(err).Msgf("failed get user entity from database")

		return err
	}

	depositEntity.Status = entity.TransactionStatusSuccess
	if err := tx.Save(&depositEntity).Error; err != nil {
		logger.Error().Err(err).Msgf("failed save user entity to database")

		return err
	}

	toUserEntity.Money = toUserEntity.Money.Add(depositEntity.Money)

	if err := tx.Save(&toUserEntity).Error; err != nil {
		logger.Error().Err(err).Msgf("failed save user entity to database")

		return err
	}

	tx.Commit()

	return nil
}

func (receiver *TransactionSchedulerService) geDepositUIDsToProcess(logger zerolog.Logger) (confirmedEntitiesUIDs []uuid.UUID, err error) {
	if err := receiver.db.Model(&entity.UserDeposit{}).
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
