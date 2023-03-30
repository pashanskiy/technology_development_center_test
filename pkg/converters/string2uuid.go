package converters

import (
	"github.com/gofrs/uuid/v5"
	"github.com/rs/zerolog"
)

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
