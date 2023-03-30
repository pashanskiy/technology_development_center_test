package interceptor

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/pereslava/grpc_zerolog"
	"github.com/rs/zerolog"

	"google.golang.org/grpc"
)

const HeaderUserInfo = "authorization"
const ReceiveRequest = "Receive request: %++v"

func GrpcServerWithInterceptors(logger zerolog.Logger) []grpc.ServerOption {
	unaryInterceptors := grpc.ChainUnaryInterceptor(
		grpc_zerolog.NewUnaryServerInterceptor(logger),
		grpcLogInterceptor(logger),
	)
	streamInterceptors := grpc.ChainStreamInterceptor(
		grpc_zerolog.NewStreamServerInterceptor(logger),
	)
	return []grpc.ServerOption{unaryInterceptors, streamInterceptors}
}

func grpcLogInterceptor(logger zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		logger = addUserInfoToLogger(logger, ctx)
		logger.Debug().Msgf(ReceiveRequest, req)

		newCtx := logger.WithContext(ctx)
		return handler(newCtx, req)
	}
}

func addUserInfoToLogger(logger zerolog.Logger, ctx context.Context) zerolog.Logger {
	headerValue := metadata.ExtractIncoming(ctx).Get(HeaderUserInfo)
	logger.Trace().Msgf("headerValue: %s", headerValue)

	if headerValue == "" {
		logger.Trace().Msg("user info was not found")

		return logger
	}

	return logger.With().Str(HeaderUserInfo, headerValue).Logger()
}
