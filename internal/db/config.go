package db

type Config struct {
	DBHost string `env:"DB_HOST" envDefault:"0.0.0.0"`
	DBPort uint16 `env:"DB_PORT" envDefault:"5432"`

	DBUser string `env:"DB_USER" envDefault:"kek"`
	DBPass string `env:"DB_PASS" envDefault:"kek"`
	DBName string `env:"DB_NAME" envDefault:"tdc"`

	MigrationsDirPath string `env:"DB_MIGRATIONS_DIR" envDefault:"./internal/db/migrations"`
}
