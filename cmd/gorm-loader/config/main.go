package main

import (
	"calamity-ai/internal"
	"fmt"
	"github.com/goccy/go-json"
	"log"
)

type config struct {
	DatabaseUrl string `env:"DATABASE_URL" validate:"required" json:"databaseUrl,omitempty"`
}

func main() {
	parsed := internal.ParseConfig[config]()

	marshal, err := json.Marshal(parsed)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(marshal))
}
