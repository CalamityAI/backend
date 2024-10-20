package internal

type dbConfig struct {
	DatabaseDsn string `env:"DATABASE_DSN" validate:"required"`
}

var DatabaseConfig = ParseConfig[dbConfig]()
