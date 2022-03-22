package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"tournament_gateway/app/api/controllers"
	"tournament_gateway/app/api/controllers/current"
	groupsControllers "tournament_gateway/app/api/controllers/groups"
	leaguesControllers "tournament_gateway/app/api/controllers/leagues"
	matchesControllers "tournament_gateway/app/api/controllers/matches"
	seasonsControllers "tournament_gateway/app/api/controllers/seasons"
	teamsControllers "tournament_gateway/app/api/controllers/teams"
	tournamentsControllers "tournament_gateway/app/api/controllers/tournaments"
	groupsService "tournament_gateway/services/groups"
	leaguesService "tournament_gateway/services/leagues"
	matchesService "tournament_gateway/services/matches"
	playersService "tournament_gateway/services/players"
	seasonsService "tournament_gateway/services/seasons"
	teamsService "tournament_gateway/services/teams"
	tournamentsService "tournament_gateway/services/tournaments"
)

type RestAPI struct {
	tournamentsService *tournamentsService.Service
	seasonsService     *seasonsService.Service
	leaguesService     *leaguesService.Service
	groupsService      *groupsService.Service
	teamsService       *teamsService.Service
	playersService     *playersService.Service
	matchesService     *matchesService.Service
	logger             *zerolog.Logger
}

func NewRestAPI(tournamentsS *tournamentsService.Service, seasonsS *seasonsService.Service,
	leaguesS *leaguesService.Service, groupsS *groupsService.Service,
	teamsS *teamsService.Service, playersS *playersService.Service,
	matchesS *matchesService.Service, logger *zerolog.Logger) *RestAPI {
	return &RestAPI{tournamentsService: tournamentsS, seasonsService: seasonsS,
		leaguesService: leaguesS, groupsService: groupsS, teamsService: teamsS,
		playersService: playersS, matchesService: matchesS, logger: logger}
}

func (a *RestAPI) RunHTTPServer(ctx context.Context) error {
	r := gin.Default()
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.Writer.Header().Set("Content-Type", "application/json")
	}))

	a.openapiHandlers(r)
	a.currentHandlers(r)
	a.tournamentsHandlers(r)
	a.seasonsHandlers(r)
	a.leaguesHandlers(r)
	a.groupsHandlers(r)
	a.teamsHandlers(r)
	a.matchesHandlers(r)

	return r.Run()
}

func (a *RestAPI) openapiHandlers(r *gin.Engine) {
	r.GET("/openapi.json", func(c *gin.Context) { http.ServeFile(c.Writer, c.Request, "./openapi.json") })
}

func (a *RestAPI) currentHandlers(r *gin.Engine) {
	getCurrents := current.NewCurrent(a.groupsService, a.seasonsService, a.tournamentsService, a.logger)
	r.GET("/api/current", getCurrents.HTTPHandler)
}

func (a *RestAPI) tournamentsHandlers(r *gin.Engine) {
	getTournaments := tournamentsControllers.NewGetTournaments(a.tournamentsService, a.logger)
	r.GET("/api/tournaments", getTournaments.HTTPHandler)

	getTournament := tournamentsControllers.NewGetTournament(a.tournamentsService, a.logger)
	r.GET("/api/tournaments/:tournament_alias", getTournament.HTTPHandler)

	updateTournament := tournamentsControllers.NewUpdateTournament(a.tournamentsService, a.logger)
	r.PUT("/api/tournaments/:tournament_alias", updateTournament.HTTPHandler)

	createTournament := tournamentsControllers.NewCreateTournament(a.tournamentsService, a.logger)
	r.POST("/api/tournaments", createTournament.HTTPHandler)

	r.OPTIONS("/api/tournaments/:tournament_alias", controllers.OptionsHTTPHandler)
	r.OPTIONS("/api/tournaments", controllers.OptionsHTTPHandler)
}

func (a *RestAPI) seasonsHandlers(r *gin.Engine) {
	getSeasons := seasonsControllers.NewGetSeasons(a.seasonsService, a.logger)
	r.GET("/api/seasons", getSeasons.HTTPHandler)

	getSeason := seasonsControllers.NewGetSeason(a.seasonsService, a.logger)
	r.GET("/api/seasons/:season_alias", getSeason.HTTPHandler)

	createSeason := seasonsControllers.NewCreateSeason(a.seasonsService, a.logger)
	r.POST("/api/seasons", createSeason.HTTPHandler)

	updateSeason := seasonsControllers.NewUpdateSeason(a.seasonsService, a.logger)
	r.PUT("/api/seasons/:season_alias", updateSeason.HTTPHandler)

	getStages := seasonsControllers.NewGetStages(a.seasonsService, a.logger)
	r.GET("/api/seasons/:season_alias/stages", getStages.HTTPHandler)

	r.OPTIONS("/api/seasons", controllers.OptionsHTTPHandler)
	r.OPTIONS("/api/seasons/:season_alias", controllers.OptionsHTTPHandler)
}

func (a *RestAPI) leaguesHandlers(r *gin.Engine) {
	findLeagues := leaguesControllers.NewFindLeagues(a.groupsService, a.logger)
	r.GET("/api/leagues/list/:tournament_alias/:season_alias/:stage_alias", findLeagues.HTTPHandler)

	getLeagues := leaguesControllers.NewGetLeagues(a.leaguesService, a.logger)
	r.GET("/api/leagues", getLeagues.HTTPHandler)

	getLeague := leaguesControllers.NewGetLeague(a.leaguesService, a.logger)
	r.GET("/api/leagues/:league_alias", getLeague.HTTPHandler)

	createLeague := leaguesControllers.NewCreateLeague(a.leaguesService, a.logger)
	r.POST("/api/leagues", createLeague.HTTPHandler)

	updateLeague := leaguesControllers.NewUpdateLeague(a.leaguesService, a.logger)
	r.PUT("/api/leagues/:league_alias", updateLeague.HTTPHandler)

	r.OPTIONS("/api/leagues", controllers.OptionsHTTPHandler)
	r.OPTIONS("/api/leagues/:league_alias", controllers.OptionsHTTPHandler)
}

func (a *RestAPI) groupsHandlers(r *gin.Engine) {
	getTeams := groupsControllers.NewGetTeams(a.groupsService, a.logger)
	r.GET("/api/groups/:group_alias/teams", getTeams.HTTPHandler)

	getMatches := groupsControllers.NewGetMatches(a.groupsService, a.logger)
	r.GET("/api/groups/:group_alias/matches", getMatches.HTTPHandler)

	getResults := groupsControllers.NewGetGroupResults(a.groupsService, a.logger)
	r.GET("/api/groups/:group_alias/results", getResults.HTTPHandler)

	findGroups := groupsControllers.NewFindGroups(a.groupsService, a.logger)
	r.GET("/api/groups/list/:tournament_alias/:season_alias/:stage_alias/:league_alias", findGroups.HTTPHandler)

	getGroups := groupsControllers.NewGetGroups(a.groupsService, a.logger)
	r.GET("/api/groups", getGroups.HTTPHandler)

	getGroup := groupsControllers.NewGetGroup(a.groupsService, a.logger)
	r.GET("/api/groups/:group_alias", getGroup.HTTPHandler)

	createGroup := groupsControllers.NewCreateGroup(a.groupsService, a.tournamentsService, a.seasonsService, a.leaguesService, a.logger)
	r.POST("/api/groups", createGroup.HTTPHandler)

	updateGroup := groupsControllers.NewUpdateGroup(a.groupsService, a.logger)
	r.PUT("/api/groups/:group_alias", updateGroup.HTTPHandler)

	r.OPTIONS("/api/groups", controllers.OptionsHTTPHandler)
	r.OPTIONS("/api/groups/:group_alias", controllers.OptionsHTTPHandler)
}

func (a *RestAPI) teamsHandlers(r *gin.Engine) {
	getTeam := teamsControllers.NewGetTeam(a.teamsService, a.logger)
	r.GET("/api/teams/:team_id", getTeam.HTTPHandler)

	getGroups := teamsControllers.NewGetGroups(a.teamsService, a.logger)
	r.GET("/api/teams/:team_id/groups", getGroups.HTTPHandler)

	getPlayers := teamsControllers.NewGetPlayers(a.playersService, a.logger)
	r.GET("/api/teams/:team_id/groups/:group_alias/players", getPlayers.HTTPHandler)
}

func (a *RestAPI) matchesHandlers(r *gin.Engine) {
	getMatch := matchesControllers.NewGetMatch(a.matchesService, a.logger)
	r.GET("/api/matches/:match_id", getMatch.HTTPHandler)
}
