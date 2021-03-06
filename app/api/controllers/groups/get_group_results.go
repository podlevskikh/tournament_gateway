package groups

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
)

type GetGroupResults struct {
	service GroupService
	logger  *zerolog.Logger
}

func NewGetGroupResults(service GroupService, logger *zerolog.Logger) *GetGroupResults {
	return &GetGroupResults{service: service, logger: logger}
}

func (s *GetGroupResults) HTTPHandler(c *gin.Context) {
	groupAlias := c.Param("group_alias")

	results, err := s.service.GetGroupResults(c.Request.Context(), groupAlias)
	if err != nil {
		s.logger.Err(err).Msg("get group results")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}

	response_factory.ReturnSuccessList(c, response_success.FromGroupResultsResponse(results), len(results))
}
