package internal

type dbConfig struct {
	DatabaseUrl string `env:"DATABASE_URL" validate:"required"`
}

var DatabaseConfig = ParseConfig[dbConfig]()
