package main

import (
	"Kitchenizer/cmd/app"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	if err := app.Run(); err != nil {
		log.Error().Err(err)
		os.Exit(1)
	}
}
