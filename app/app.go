package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"vollyemsk_tournament_gateway/resources"
)

type App struct {
	res    *resources.Resources
	logger *log.Logger
}

func (a *App) Start(ctx context.Context, logger *log.Logger) error {
	a.res = resources.Get(ctx, logger)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
