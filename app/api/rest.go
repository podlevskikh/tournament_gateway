package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	groupsControllers "vollyemsk_tournament_gateway/app/api/controllers/groups"
	leaguesControllers "vollyemsk_tournament_gateway/app/api/controllers/leagues"
	seasonsControllers "vollyemsk_tournament_gateway/app/api/controllers/seasons"
	tournamentsControllers "vollyemsk_tournament_gateway/app/api/controllers/tournaments"
	groupsService "vollyemsk_tournament_gateway/services/groups"
	leaguesService "vollyemsk_tournament_gateway/services/leagues"
	seasonsService "vollyemsk_tournament_gateway/services/seasons"
	tournamentsService "vollyemsk_tournament_gateway/services/tournaments"
)

type RestAPI struct {
	tournamentsService *tournamentsService.Service
	seasonsService     *seasonsService.Service
	leaguesService     *leaguesService.Service
	groupsService      *groupsService.Service
	logger             *zerolog.Logger
}

func NewRestAPI(tournamentsS *tournamentsService.Service, seasonsS *seasonsService.Service,
	leaguesS *leaguesService.Service, groupsS *groupsService.Service, logger *zerolog.Logger) *RestAPI {
	return &RestAPI{tournamentsService: tournamentsS, seasonsService: seasonsS,
		leaguesService: leaguesS, groupsService: groupsS, logger: logger}
}

func (a *RestAPI) RunHTTPServer(ctx context.Context) error {
	r := gin.Default()
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.Writer.Header().Set("Content-Type", "application/json")
	}))

	a.tournamentsHandlers(r)
	a.seasonsHandlers(r)
	a.leaguesHandlers(r)
	a.groupsHandlers(r)

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

func (a *RestAPI) leaguesHandlers(r *gin.Engine) {
	getLeagues := leaguesControllers.NewGetLeagues(a.groupsService, a.logger)
	r.GET("/api/leagues/:tournament_alias/:season_alias/:stage_alias", getLeagues.HTTPHandler)
}

func (a *RestAPI) groupsHandlers(r *gin.Engine) {
	getTeams := groupsControllers.NewGetTeams(a.groupsService, a.logger)
	r.GET("/api/groups/teams/:group_alias", getTeams.HTTPHandler)

	getMatches := groupsControllers.NewGetMatches(a.groupsService, a.logger)
	r.GET("/api/groups/matches/:group_alias", getMatches.HTTPHandler)

	getGroups := groupsControllers.NewGetGroups(a.groupsService, a.logger)
	r.GET("/api/groups/:tournament_alias/:season_alias/:stage_alias/:league_alias", getGroups.HTTPHandler)
}
