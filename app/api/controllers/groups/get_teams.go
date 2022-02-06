package groups

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetTeams struct {
	service Service
	logger  *zerolog.Logger
}

func NewGetTeams(service Service, logger *zerolog.Logger) *GetTeams {
	return &GetTeams{service: service, logger: logger}
}

func (s *GetTeams) HTTPHandler(c *gin.Context) {
	groupAlias := c.Param("group_alias")

	teams, err := s.service.GetGroupTeams(c.Request.Context(), groupAlias)
	if err != nil {
		s.logger.Err(err).Msg("get group teams")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}

	response_factory.ReturnSuccess(c, response_success.FromTeamsResponse(teams))
}
