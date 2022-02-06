package teams

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"strconv"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetTeam struct {
	service TeamsService
	logger  *zerolog.Logger
}

func NewGetTeam(service TeamsService, logger *zerolog.Logger) *GetTeam {
	return &GetTeam{service: service, logger: logger}
}

func (s *GetTeam) HTTPHandler(c *gin.Context) {
	teamID, err := strconv.Atoi(c.Param("team_id"))
	if err != nil {
		s.logger.Err(err).Msg("get team request")
		response_factory.ReturnError(c, response_error.ParseRequest)
		return
	}

	team, err := s.service.GetTeam(c.Request.Context(), teamID)
	if err != nil {
		s.logger.Err(err).Msg("get team")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}

	response_factory.ReturnSuccess(c, response_success.FromTeamResponse(team))
}
