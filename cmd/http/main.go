package main

import (
	. "calamity-ai/internal/http/infra"
)

func main() {

	router := SetupRouter()

	router.Run()
}
