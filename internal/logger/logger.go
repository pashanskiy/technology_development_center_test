package logger

import (
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ParseEnvLoggerEnv(value string) zerolog.Level {
	v := strings.ToLower(value)

	var level zerolog.Level

	switch v {
	case "info":
		level = zerolog.InfoLevel
	case "warn":
		level = zerolog.WarnLevel
	case "error":
		level = zerolog.ErrorLevel
	case "debug":
		level = zerolog.DebugLevel
	case "trace":
		level = zerolog.TraceLevel
	default:
		log.Warn().Msgf("Unknown logging level: %s", value)

		level = zerolog.InfoLevel
	}

	return level
}

func InitRootLogger(globalLevel zerolog.Level, serviceName, buildDate string) zerolog.Logger {
	zerolog.SetGlobalLevel(globalLevel)

	log.Logger = log.With().
		Str("service", serviceName).
		Str("buildDate", buildDate).
		Caller().
		Logger()
	zerolog.DefaultContextLogger = &log.Logger

	return log.Logger
}
