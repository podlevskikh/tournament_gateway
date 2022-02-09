package teams

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"strconv"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetPlayers struct {
	service PlayersService
	logger  *zerolog.Logger
}

func NewGetPlayers(service PlayersService, logger *zerolog.Logger) *GetPlayers {
	return &GetPlayers{service: service, logger: logger}
}

func (s *GetPlayers) HTTPHandler(c *gin.Context) {
	teamID, err := strconv.Atoi(c.Param("team_id"))
	if err != nil {
		s.logger.Err(err).Msg("get team request")
		response_factory.ReturnError(c, response_error.ParseRequest)
		return
	}

	groupAlias := c.Param("group_alias")

	players, err := s.service.GetPlayersByTeamAndGroup(c.Request.Context(), teamID, groupAlias)
	if err != nil {
		s.logger.Err(err).Msg("get team")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}

	response_factory.ReturnSuccessList(c, response_success.FromPlayersResponse(players), len(players))
}
