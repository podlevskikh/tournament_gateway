package seasons

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetStages struct {
	service Service
	logger  *zerolog.Logger
}

func NewGetStages(service Service, logger *zerolog.Logger) *GetStages {
	return &GetStages{
		service: service,
		logger:  logger,
	}
}

func (s *GetStages) HTTPHandler(c *gin.Context) {
	seasonAlias := c.Param("season_alias")

	stages, err := s.service.GetStagesBySeasonAlias(c.Request.Context(), seasonAlias)
	if err != nil {
		s.logger.Err(err).Msg("get stages")
		//todo season not found
		response_factory.ReturnError(c, response_error.Internal)
		return
	}
	response_factory.ReturnSuccessList(c, response_success.FromStagesResponse(stages), len(stages))
}

