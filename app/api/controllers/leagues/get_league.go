package leagues

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetLeague struct {
	service LeagueService
	logger  *zerolog.Logger
}

func NewGetLeague(service LeagueService, logger *zerolog.Logger) *GetLeague {
	return &GetLeague{
		service: service,
		logger:  logger,
	}
}

func (s *GetLeague) HTTPHandler(c *gin.Context) {
	leagueAlias := c.Param("league_alias")

	l, err := s.service.GetLeague(c.Request.Context(), leagueAlias)
	if err != nil {
		s.logger.Err(err).Msg("get league")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}
	response_factory.ReturnSuccess(c, response_success.FromLeagueResponse(l))
}
