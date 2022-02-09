package seasons

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetSeason struct {
	service Service
	logger  *zerolog.Logger
}

func NewGetSeason(service Service, logger *zerolog.Logger) *GetSeason {
	return &GetSeason{
		service: service,
		logger:  logger,
	}
}

func (s *GetSeason) HTTPHandler(c *gin.Context) {
	seasonAlias := c.Param("season_alias")

	season, err := s.service.GetSeason(c.Request.Context(), seasonAlias)
	if err != nil {
		s.logger.Err(err).Msg("get season")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}
	response_factory.ReturnSuccess(c, response_success.FromSeasonResponse(season))
}
