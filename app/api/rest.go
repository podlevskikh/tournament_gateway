package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	seasonsControllers "vollyemsk_tournament_gateway/app/api/controllers/seasons"
	tournamentsControllers "vollyemsk_tournament_gateway/app/api/controllers/tournaments"
	seasonsService "vollyemsk_tournament_gateway/services/seasons"
	tournamentsService "vollyemsk_tournament_gateway/services/tournaments"
)

type RestAPI struct {
	tournamentsService *tournamentsService.Service
	seasonsService     *seasonsService.Service
	logger             *zerolog.Logger
}

func NewRestAPI(tournamentsS *tournamentsService.Service, seasonsS *seasonsService.Service, logger *zerolog.Logger) *RestAPI {
	return &RestAPI{tournamentsService: tournamentsS, seasonsService: seasonsS, logger: logger}
}

func (a *RestAPI) RunHTTPServer(ctx context.Context) error {
	r := gin.Default()
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.Writer.Header().Set("Content-Type", "application/json")
	}))

	a.tournamentsHandlers(r)
	a.seasonsHandlers(r)

	return r.Run()
}

func (a *RestAPI) tournamentsHandlers(r *gin.Engine) {
	getTournaments := tournamentsControllers.NewGetTournaments(a.tournamentsService, a.logger)
	r.GET("/api/tournaments", getTournaments.HTTPHandler)
}

func (a *RestAPI) seasonsHandlers(r *gin.Engine) {
	getSeasons := seasonsControllers.NewGetSeasons(a.seasonsService, a.logger)
	r.GET("/api/seasons", getSeasons.HTTPHandler)

	getStages := seasonsControllers.NewGetStages(a.seasonsService, a.logger)
	r.GET("/api/seasons/:season_alias/stages", getStages.HTTPHandler)
}