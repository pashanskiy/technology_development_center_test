package main

import (
	config "technology_development_center_test/config"
	dbService "technology_development_center_test/internal/db"
	envParse "technology_development_center_test/internal/env-parse"
	loggerService "technology_development_center_test/internal/logger"

	"github.com/rs/zerolog/log"
)

// defaults for -ldflags
var ServiceName = "Storage"
var BuildDate = "nil"

func main() {
	var serviceConfig config.Config
	envParse.ReadEnvConfig(log.Logger, &serviceConfig)

	logger := loggerService.InitRootLogger(
		loggerService.ParseEnvLoggerEnv(serviceConfig.LogRootLevel),
		ServiceName,
		BuildDate,
	)

	db := dbService.NewDB(logger).Gorm

	_ = db

}
