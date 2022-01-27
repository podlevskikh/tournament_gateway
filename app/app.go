package app

import (
	"context"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"vollyemsk_tournament_gateway/app/api"
	"vollyemsk_tournament_gateway/resources"
	groupsService "vollyemsk_tournament_gateway/services/groups"
	groupsRepository "vollyemsk_tournament_gateway/services/groups/repository"
	leaguesService "vollyemsk_tournament_gateway/services/leagues"
	leaguesRepository "vollyemsk_tournament_gateway/services/leagues/repository"
	seasonsService "vollyemsk_tournament_gateway/services/seasons"
	seasonsRepository "vollyemsk_tournament_gateway/services/seasons/repository"
	tournamentsService "vollyemsk_tournament_gateway/services/tournaments"
	tournamentsRepository "vollyemsk_tournament_gateway/services/tournaments/repository"
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

	restAPI := api.NewRestAPI(repositoryService, seasonService, leagueService, groupService, a.logger)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return restAPI.RunHTTPServer(ctx)
	})
	return g.Wait()
}
