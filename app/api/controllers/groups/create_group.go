package groups

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/controllers/groups/requests"
	"tournament_gateway/app/api/request_factory"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
	"tournament_gateway/models/groups"
	"tournament_gateway/models/seasons"
)

type CreateGroup struct {
	groupService      GroupService
	tournamentService TournamentService
	seasonService     SeasonService
	leagueService     LeagueService
	logger            *zerolog.Logger
}

func NewCreateGroup(groupService GroupService,
	tournamentService TournamentService,
	seasonService SeasonService,
	leagueService LeagueService,
	logger *zerolog.Logger) *CreateGroup {
	return &CreateGroup{
		groupService:      groupService,
		tournamentService: tournamentService,
		seasonService:     seasonService,
		leagueService:     leagueService,
		logger:            logger,
	}
}

func (s *CreateGroup) HTTPHandler(c *gin.Context) {
	createRequest := requests.GroupRequest{}
	err := request_factory.ReadJSONRequestBody(c, &createRequest)
	if err != nil {
		s.logger.Err(err).Msg("create group request")
		return
	}
	_, err = s.tournamentService.GetTournament(c.Request.Context(), createRequest.TournamentAlias)
	if err != nil {
		s.logger.Err(err).Msg("tournament not found")
		response_factory.ReturnError(c, response_error.TournamentNotFound)
		return
	}

	se, err := s.seasonService.GetSeason(c.Request.Context(), createRequest.SeasonAlias)
	if err != nil {
		s.logger.Err(err).Msg("season not found")
		response_factory.ReturnError(c, response_error.SeasonNotFound)
		return
	}

	if !hasStage(se, createRequest.StageAlias) {
		s.logger.Err(err).Msg("stage not found")
		response_factory.ReturnError(c, response_error.StageNotFound)
		return
	}

	_, err = s.leagueService.GetLeague(c.Request.Context(), createRequest.LeagueAlias)
	if err != nil {
		s.logger.Err(err).Msg("league not found")
		response_factory.ReturnError(c, response_error.LeagueNotFound)
		return
	}

	g := groups.Group{
		Alias:           createRequest.Alias,
		ShortName:       createRequest.ShortName,
		Name:            createRequest.Name,
		Description:     createRequest.Description,
		TournamentAlias: createRequest.TournamentAlias,
		SeasonAlias:     createRequest.SeasonAlias,
		StageAlias:      createRequest.StageAlias,
		LeagueAlias:     createRequest.LeagueAlias,
	}

	gr, err := s.groupService.CreateGroup(c.Request.Context(), g)
	if err != nil {
		s.logger.Err(err).Msg("create group")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}
	response_factory.ReturnSuccess(c, response_success.FromGroupResponse(gr))
}

func hasStage(se *seasons.Season, stageAlias string) bool {
	for _, st := range se.Stages {
		if st.Alias == stageAlias {
			return true
		}
	}
	return false
}
