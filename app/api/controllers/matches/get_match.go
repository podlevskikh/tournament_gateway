package matches

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"strconv"
	"vollyemsk_tournament_gateway/app/api/response_error"
	"vollyemsk_tournament_gateway/app/api/response_factory"
	"vollyemsk_tournament_gateway/app/api/response_success"
)

type GetMatch struct {
	service Service
	logger  *zerolog.Logger
}

func NewGetMatch(service Service, logger *zerolog.Logger) *GetMatch {
	return &GetMatch{service: service, logger: logger}
}

func (s *GetMatch) HTTPHandler(c *gin.Context) {
	matchID, err := strconv.Atoi(c.Param("match_id"))
	if err != nil {
		s.logger.Err(err).Msg("get match request")
		response_factory.ReturnError(c, response_error.ParseRequest)
		return
	}

	match, err := s.service.GetMatch(c.Request.Context(), matchID)
	if err != nil {
		s.logger.Err(err).Msg("get leagues")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}

	response_factory.ReturnSuccess(c, response_success.FromMatchResponse(match))
}
