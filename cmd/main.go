package main

import (
	"context"
	"github.com/rs/zerolog"
	"os"
	"tournament_gateway/app"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger().With().Str("source", "tournament_gateway").Logger()
	ctx := context.Background()
	a := app.App{}
	if err := a.Start(ctx, &logger); err != nil {
		logger.Fatal().Err(err).Msg("server start")
	}
}