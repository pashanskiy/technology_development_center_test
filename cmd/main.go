package main

import (
	"fmt"
	"net"
	apiTransactionsService "technology_development_center_test/api/transactions"
	apiUserService "technology_development_center_test/api/user"
	config "technology_development_center_test/config"
	dbService "technology_development_center_test/internal/db"
	envParse "technology_development_center_test/internal/env-parse"
	grpcInterceptors "technology_development_center_test/internal/interceptor"
	loggerService "technology_development_center_test/internal/logger"
	transactionsService "technology_development_center_test/internal/transactions-service"
	userService "technology_development_center_test/internal/user-service"
	schedulerService "technology_development_center_test/internal/transactions-service/scheduler"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	grpcServer := grpc.NewServer(grpcInterceptors.GrpcServerWithInterceptors(logger)...)

	apiUserService.RegisterUserServiceServer(grpcServer, userService.NewUserService(db))
	apiTransactionsService.RegisterTransactionsServiceServer(grpcServer, transactionsService.NewTransactionsService(db))

	scheduler := schedulerService.NewTransactionSchedulerService(db, &logger)
	scheduler.RunSchedule(logger)

	runService(serviceConfig, grpcServer)
}

func runService(serviceConfig config.Config, grpcServer *grpc.Server) {
	if serviceConfig.GrpcReflectionEnabled {
		reflection.Register(grpcServer)
	}

	listen, err := net.Listen(
		"tcp4",
		fmt.Sprintf("%s:%d", serviceConfig.GrpcServerHost, serviceConfig.GrpcServerPort),
	)
	if err != nil {
		log.Panic().Err(err).Msg("")
	}

	log.Info().Msgf("gRPC listen: %s", listen.Addr().String())

	if err = grpcServer.Serve(listen); err != nil {
		log.Panic().Err(err).Msg("")
	}
}
