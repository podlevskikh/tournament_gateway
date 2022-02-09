package groups

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetGroups struct {
	service Service
	logger  *zerolog.Logger
}

func NewGetGroups(service Service, logger *zerolog.Logger) *GetGroups {
	return &GetGroups{service: service, logger: logger}
}

func (s *GetGroups) HTTPHandler(c *gin.Context) {
	tournamentAlias := c.Param("tournament_alias")
	seasonAlias := c.Param("season_alias")
	stageAlias := c.Param("stage_alias")
	leagueAlias := c.Param("league_alias")

	groups, err := s.service.GetGroupsByTournamentSeasonStageLeagues(c.Request.Context(), tournamentAlias, seasonAlias, stageAlias, leagueAlias)
	if err != nil {
		s.logger.Err(err).Msg("get groups")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}

	response_factory.ReturnSuccessList(c, response_success.FromGroupsResponse(groups), len(groups))
}
