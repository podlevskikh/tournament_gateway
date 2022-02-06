package app

import (
	"context"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"tournament_gateway/app/api"
	"tournament_gateway/resources"
	groupsService "tournament_gateway/services/groups"
	groupsRepository "tournament_gateway/services/groups/repository"
	leaguesService "tournament_gateway/services/leagues"
	leaguesRepository "tournament_gateway/services/leagues/repository"
	matchesService "tournament_gateway/services/matches"
	matchesRepository "tournament_gateway/services/matches/repository"
	playersService "tournament_gateway/services/players"
	playersRepository "tournament_gateway/services/players/repository"
	seasonsService "tournament_gateway/services/seasons"
	seasonsRepository "tournament_gateway/services/seasons/repository"
	teamsService "tournament_gateway/services/teams"
	teamsRepository "tournament_gateway/services/teams/repository"
	tournamentsService "tournament_gateway/services/tournaments"
	tournamentsRepository "tournament_gateway/services/tournaments/repository"
)

type App struct {
	res    *resources.Resources
	logger *zerolog.Logger
}

func (a *App) Start(ctx context.Context, logger *zerolog.Logger) error {
	a.logger = logger
	a.res = resources.Get(ctx, logger)

	repositoryService := tournamentsService.NewService(tournamentsRepository.NewDb(a.res.DB))
	seasonService := seasonsService.NewService(seasonsRepository.NewDb(a.res.DB))
	leagueService := leaguesService.NewService(leaguesRepository.NewDb(a.res.DB))
	groupService := groupsService.NewService(groupsRepository.NewDb(a.res.DB))
	teamService := teamsService.NewService(teamsRepository.NewDb(a.res.DB))
	playerService := playersService.NewService(playersRepository.NewDb(a.res.DB))
	matchService := matchesService.NewService(matchesRepository.NewDb(a.res.DB))

	restAPI := api.NewRestAPI(repositoryService, seasonService, leagueService, groupService, teamService, playerService, matchService, a.logger)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return restAPI.RunHTTPServer(ctx)
	})
	return g.Wait()
}
