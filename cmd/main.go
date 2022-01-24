package main

import (
	"context"
	"github.com/rs/zerolog"
	"vollyemsk_tournament_gateway/app"
)

var Logger zerolog.Logger

func main() {
	logger := Logger.With().Str("source", "tournament_gateway").Logger()
	ctx := context.Background()
	a := app.App{}
	if err := a.Start(ctx, &logger); err != nil {
		logger.Fatal().Err(err).Msg("server start")
	}
}