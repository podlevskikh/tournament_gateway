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
)

type UpdateGroup struct {
	service GroupService
	logger  *zerolog.Logger
}

func NewUpdateGroup(service GroupService, logger *zerolog.Logger) *UpdateGroup {
	return &UpdateGroup{service: service, logger: logger}
}

func (s *UpdateGroup) HTTPHandler(c *gin.Context) {
	createRequest := requests.GroupRequest{}
	err := request_factory.ReadJSONRequestBody(c, &createRequest)
	if err != nil {
		s.logger.Err(err).Msg("update group request")
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

	gr, err := s.service.UpdateGroup(c.Request.Context(), g)
	if err != nil {
		s.logger.Err(err).Msg("update group")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}
	response_factory.ReturnSuccess(c, response_success.FromGroupResponse(gr, true))
}
