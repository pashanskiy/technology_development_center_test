package config

type Config struct {
	GrpcServerHost        string `env:"GRPC_SERVER_HOST" envDefault:"0.0.0.0"`
	GrpcServerPort        int    `env:"GRPC_SERVER_PORT" envDefault:"9000"`
	GrpcReflectionEnabled bool   `env:"GRPC_REFLECTION_ENABLED" envDefault:"true"`

	LogRootLevel string `env:"LOG_ROOT_LEVEL" envDefault:"TRACE"`
}
