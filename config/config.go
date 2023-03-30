package config

type Config struct {
	LogRootLevel string `env:"LOG_ROOT_LEVEL" envDefault:"TRACE"`
}
