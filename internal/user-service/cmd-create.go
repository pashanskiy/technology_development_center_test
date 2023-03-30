package storage

import (
	"context"
	apiCore "technology_development_center_test/api/core"
	"technology_development_center_test/internal/db/entity"
	errService "technology_development_center_test/internal/user-service/errors"

	"github.com/rs/zerolog"
)

func (receiver *UserService) Create(ctx context.Context, request *apiCore.Empty) (*apiCore.Uid, error) {
	logger := zerolog.Ctx(ctx)

	userEntity := &entity.UserEntity{}

	if err := receiver.db.Create(userEntity).Error; err != nil {
		logger.Error().Err(err).Msg("failed create user")

		return nil, errService.ErrCreateUser
	}

	logger.Trace().Msgf("user successfully created. uid: %s", userEntity.UID.String())

	return &apiCore.Uid{Uid: userEntity.UID.String()}, nil
}
