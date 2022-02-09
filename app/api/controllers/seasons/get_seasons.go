package seasons

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetSeasons struct {
	service Service
	logger  *zerolog.Logger
}

func NewGetSeasons(service Service, logger *zerolog.Logger) *GetSeasons {
	return &GetSeasons{
		service: service,
		logger:  logger,
	}
}

func (s *GetSeasons) HTTPHandler(c *gin.Context) {
	seasons, err := s.service.GetSeasons(c.Request.Context())
	if err != nil {
		s.logger.Err(err).Msg("get seasons")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}
	response_factory.ReturnSuccessList(c, response_success.FromSeasonsResponse(seasons), len(seasons))
}
