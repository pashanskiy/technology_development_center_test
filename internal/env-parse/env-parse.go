package envParse

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
)

func ReadEnvConfig(logger zerolog.Logger, cfg interface{}, opts ...env.Options) {
	if err := env.Parse(cfg, opts...); err != nil {
		logger.Panic().Err(err).Msg("Failed to read configuration from environment")
	}

	logger.Trace().Msgf("cfg: %++v", cfg)
}
