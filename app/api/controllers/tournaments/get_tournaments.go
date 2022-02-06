package tournaments

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetTournaments struct {
	service Service
	logger  *zerolog.Logger
}

func NewGetTournaments(service Service, logger *zerolog.Logger) *GetTournaments {
	return &GetTournaments{service: service, logger: logger}
}

func (s *GetTournaments) HTTPHandler(c *gin.Context) {
	t, err := s.service.GetTournaments(c.Request.Context())
	if err != nil {
		s.logger.Err(err).Msg("get tournaments")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}
	response_factory.ReturnSuccess(c, response_success.FromTournamentResponse(t))
}
