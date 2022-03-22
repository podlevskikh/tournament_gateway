package current

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type Current struct {
	groupService      GroupService
	seasonService     SeasonService
	tournamentService TournamentService
	logger            *zerolog.Logger
}

func NewCurrent(groupService GroupService, seasonService SeasonService, tournamentService TournamentService, logger *zerolog.Logger) *Current {
	return &Current{
		groupService:      groupService,
		seasonService:     seasonService,
		tournamentService: tournamentService,
		logger:            logger,
	}
}

func (s *Current) HTTPHandler(c *gin.Context) {
	gs, err := s.groupService.GetGroupsByTournamentSeasonStageLeagues(c.Request.Context(), "", "", "", "")
	if err != nil {
		s.logger.Err(err).Msg("get groups")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}

	response_factory.ReturnSuccess(c, response_success.FromCurrentResponse(gs))
}