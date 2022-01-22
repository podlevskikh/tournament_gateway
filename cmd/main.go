package main

import (
	"context"
	"log"
	"vollyemsk_tournament_gateway/app"
)

func main() {
	logger := log.Default()
	ctx := context.Background()
	a := app.App{}
	if err := a.Start(ctx, logger); err != nil {
		logger.Fatal(err)
	}
}