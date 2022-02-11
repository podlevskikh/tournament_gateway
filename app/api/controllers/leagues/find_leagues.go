package leagues

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type FindLeagues struct {
	service GroupService
	logger  *zerolog.Logger
}

func NewFindLeagues(service GroupService, logger *zerolog.Logger) *FindLeagues {
	return &FindLeagues{service: service, logger: logger}
}

func (s *FindLeagues) HTTPHandler(c *gin.Context) {
	tournamentAlias := c.Param("tournament_alias")
	seasonAlias := c.Param("season_alias")
	stageAlias := c.Param("stage_alias")

	leagues, err := s.service.GetLeaguesByTournamentSeasonStage(c.Request.Context(), tournamentAlias, seasonAlias, stageAlias)
	if err != nil {
		s.logger.Err(err).Msg("find leagues")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}

	response_factory.ReturnSuccessList(c, response_success.FromLeaguesResponse(leagues), len(leagues))
}
