package groups

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetGroups struct {
	service GroupService
	logger  *zerolog.Logger
}

func NewGetGroups(service GroupService, logger *zerolog.Logger) *GetGroups {
	return &GetGroups{service: service, logger: logger}
}

func (s *GetGroups) HTTPHandler(c *gin.Context) {
	groups, err := s.service.GetGroups(c.Request.Context())
	if err != nil {
		s.logger.Err(err).Msg("get groups")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}

	response_factory.ReturnSuccessList(c, response_success.FromGroupsResponse(groups), len(groups))
}
