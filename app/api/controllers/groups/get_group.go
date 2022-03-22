package groups

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetGroup struct {
	service GroupService
	logger  *zerolog.Logger
}

func NewGetGroup(service GroupService, logger *zerolog.Logger) *GetGroup {
	return &GetGroup{
		service: service,
		logger:  logger,
	}
}

func (s *GetGroup) HTTPHandler(c *gin.Context) {
	groupAlias := c.Param("group_alias")

	g, err := s.service.GetGroup(c.Request.Context(), groupAlias)
	if err != nil {
		s.logger.Err(err).Msg("get group")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}
	response_factory.ReturnSuccess(c, response_success.FromGroupResponse(g, true))
}
