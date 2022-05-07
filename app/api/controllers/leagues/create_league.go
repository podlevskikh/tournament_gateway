package leagues

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/controllers/leagues/requests"
	"tournament_gateway/app/api/request_factory"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
	"tournament_gateway/models/leagues"
)

type CreateLeague struct {
	service LeagueService
	logger  *zerolog.Logger
}

func NewCreateLeague(service LeagueService, logger *zerolog.Logger) *CreateLeague {
	return &CreateLeague{service: service, logger: logger}
}

func (s *CreateLeague) HTTPHandler(c *gin.Context) {
	createRequest := requests.LeagueRequest{}
	err := request_factory.ReadJSONRequestBody(c, &createRequest)
	if err != nil {
		s.logger.Err(err).Msg("create league request")
		return
	}

	l := leagues.League{
		Alias:          createRequest.Alias,
		ShortName:      createRequest.ShortName,
		Name:           createRequest.Name,
		StrengthWeight: createRequest.StrengthWeight,
	}

	le, err := s.service.CreateLeague(c.Request.Context(), l)
	if err != nil {
		s.logger.Err(err).Msg("create league")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}
	response_factory.ReturnSuccess(c, response_success.FromLeagueResponse(le))
}
