package storage

import (
	"context"
	"technology_development_center_test/api/core"
	apiUserService "technology_development_center_test/api/user"
	"technology_development_center_test/internal/db/entity"
	errService "technology_development_center_test/internal/user-service/errors"
	"technology_development_center_test/pkg/converters"
	"time"

	"github.com/shopspring/decimal"

	"github.com/rs/zerolog"
)

func (receiver *UserService) Get(ctx context.Context, request *apiUserService.UserInfo) (*apiUserService.GetUsers, error) {
	logger := zerolog.Ctx(ctx)

	userUID, err := converters.String2UUID(logger, request.Uid)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert string to uuid")

		return nil, err
	}

	userEntity := &entity.UserEntity{
		UID:       userUID,
		Money:     decimal.Decimal{},
		CreatedAt: time.Time{},
	}

	userEntities := &[]entity.UserEntity{}

	if dbErr := receiver.db.Find(userEntities, userEntity).Error; dbErr != nil {
		logger.Error().Err(dbErr).Msg("failed get user")

		return nil, errService.ErrGetUser
	}

	usersList := make([]*apiUserService.UserInfo, len(*userEntities))

	for idx, user := range *userEntities {
		stringUID := user.UID.String()
		date := user.CreatedAt.UnixNano()
		usersList[idx] = &apiUserService.UserInfo{
			Uid:        &stringUID,
			Money:      &core.Money{Money: user.Money.String()},
			Registered: &date,
		}
	}

	logger.Trace().Msgf("user successfully was taken. users len: %d", len(usersList))

	return &apiUserService.GetUsers{
		UsersInfo: usersList,
	}, nil
}
