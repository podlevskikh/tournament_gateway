package groups

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"vollyemsk_tournament_gateway/app/api/response_error"
	"vollyemsk_tournament_gateway/app/api/response_factory"
	"vollyemsk_tournament_gateway/app/api/response_success"
)

type GetMatches struct {
	service Service
	logger  *zerolog.Logger
}

func NewGetMatches(service Service, logger *zerolog.Logger) *GetMatches {
	return &GetMatches{service: service, logger: logger}
}

func (s *GetMatches) HTTPHandler(c *gin.Context) {
	groupAlias := c.Param("group_alias")

	group, err := s.service.GetGroupWithMatches(c.Request.Context(), groupAlias)
	if err != nil {
		s.logger.Err(err).Msg("get groups")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}

	response_factory.ReturnSuccess(c, response_success.FromMatchesResponse(group.Matches))
}
