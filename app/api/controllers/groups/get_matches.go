package groups

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetMatches struct {
	service GroupService
	logger  *zerolog.Logger
}

func NewGetMatches(service GroupService, logger *zerolog.Logger) *GetMatches {
	return &GetMatches{service: service, logger: logger}
}

func (s *GetMatches) HTTPHandler(c *gin.Context) {
	groupAlias := c.Param("group_alias")

	matches, err := s.service.GetGroupMatches(c.Request.Context(), groupAlias)
	if err != nil {
		s.logger.Err(err).Msg("get group matches")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}

	response_factory.ReturnSuccessList(c, response_success.FromMatchesResponse(matches), len(matches))
}
