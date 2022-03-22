package teams

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"strconv"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetGroups struct {
	service TeamsService
	logger  *zerolog.Logger
}

func NewGetGroups(service TeamsService, logger *zerolog.Logger) *GetGroups {
	return &GetGroups{service: service, logger: logger}
}

func (s *GetGroups) HTTPHandler(c *gin.Context) {
	teamID, err := strconv.Atoi(c.Param("team_id"))
	if err != nil {
		s.logger.Err(err).Msg("get team request")
		response_factory.ReturnError(c, response_error.ParseRequest)
		return
	}

	team, err := s.service.GetTeamWithGroups(c.Request.Context(), teamID)
	if err != nil {
		s.logger.Err(err).Msg("get team")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}

	response_factory.ReturnSuccessList(c, response_success.FromGroupsResponse(team.Groups, true), len(team.Groups))
}
