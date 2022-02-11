package leagues

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetLeagues struct {
	service LeagueService
	logger  *zerolog.Logger
}

func NewGetLeagues(service LeagueService, logger *zerolog.Logger) *GetLeagues {
	return &GetLeagues{service: service, logger: logger}
}

func (s *GetLeagues) HTTPHandler(c *gin.Context) {
	leagues, err := s.service.GetLeagues(c.Request.Context())
	if err != nil {
		s.logger.Err(err).Msg("get leagues")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}

	response_factory.ReturnSuccessList(c, response_success.FromLeaguesResponse(leagues), len(leagues))
}
