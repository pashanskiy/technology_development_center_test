package storage

import (
	"context"
	apiCore "technology_development_center_test/api/core"
	"technology_development_center_test/internal/db/entity"
	errService "technology_development_center_test/internal/user-service/errors"
	"technology_development_center_test/pkg/converters"

	"github.com/rs/zerolog"
)

func (receiver *UserService) Delete(ctx context.Context, request *apiCore.Uid) (*apiCore.Empty, error) {
	logger := zerolog.Ctx(ctx)

	userUID, err := converters.String2UUID(logger, &request.Uid)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert string to uuid")

		return nil, err
	}

	userEntity := &entity.UserEntity{
		UID: userUID,
	}

	if dbErr := receiver.db.Delete(userEntity, userEntity).Error; dbErr != nil {
		logger.Error().Err(dbErr).Msg("failed delete user")

		return nil, errService.ErrDeleteUser
	}

	logger.Trace().Msgf("user successfully deleted. userUID: %s", userUID.String())

	return &apiCore.Empty{}, nil
}
