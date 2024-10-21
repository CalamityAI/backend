package main

import (
	"ariga.io/atlas-provider-gorm/gormschema"
	"calamity-ai/internal/http/infra"
	"fmt"
	"io"
	"os"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(&infra.Greeting{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
