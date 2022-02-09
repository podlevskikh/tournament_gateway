package tournaments

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetTournament struct {
	service Service
	logger  *zerolog.Logger
}

func NewGetTournament(service Service, logger *zerolog.Logger) *GetTournament {
	return &GetTournament{service: service, logger: logger}
}

func (s *GetTournament) HTTPHandler(c *gin.Context) {
	tournamentAlias := c.Param("tournament_alias")

	t, err := s.service.GetTournament(c.Request.Context(), tournamentAlias)
	if err != nil {
		s.logger.Err(err).Msg("get tournaments") //todo 404
		response_factory.ReturnError(c, response_error.Internal)
		return
	}
	response_factory.ReturnSuccess(c, response_success.FromTournamentResponse(t))
}
