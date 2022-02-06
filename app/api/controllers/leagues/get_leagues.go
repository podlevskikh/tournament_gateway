package leagues

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetLeagues struct {
	service Service
	logger  *zerolog.Logger
}

func NewGetLeagues(service Service, logger *zerolog.Logger) *GetLeagues {
	return &GetLeagues{service: service, logger: logger}
}

func (s *GetLeagues) HTTPHandler(c *gin.Context) {
	tournamentAlias := c.Param("tournament_alias")
	seasonAlias := c.Param("season_alias")
	stageAlias := c.Param("stage_alias")

	leagues, err := s.service.GetLeaguesByTournamentSeasonStage(c.Request.Context(), tournamentAlias, seasonAlias, stageAlias)
	if err != nil {
		s.logger.Err(err).Msg("get leagues")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}

	response_factory.ReturnSuccess(c, response_success.FromLeaguesResponse(leagues))
}
