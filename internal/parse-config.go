package internal

import (
	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
)

func ParseConfig[TConfig interface{}]() TConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	validate := validator.New()

	cfg, err := env.ParseAs[TConfig]()
	if err != nil {
		log.Fatal("Could not parse env file. Make sure the file exists")
	}

	errors := validate.Struct(cfg)
	if errors != nil {
		validationErrors := errors.(validator.ValidationErrors)
		log.Fatalln("Env file validation:\n", validationErrors)
	}

	return cfg
}
