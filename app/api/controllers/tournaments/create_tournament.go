package tournaments

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/controllers/tournaments/requests"
	"tournament_gateway/app/api/request_factory"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
	"tournament_gateway/models/tournaments"
)

type CreateTournament struct {
	service Service
	logger  *zerolog.Logger
}

func NewCreateTournament(service Service, logger *zerolog.Logger) *CreateTournament {
	return &CreateTournament{service: service, logger: logger}
}

func (s *CreateTournament) HTTPHandler(c *gin.Context) {
	createRequest := requests.TournamentRequest{}
	err := request_factory.ReadJSONRequestBody(c, &createRequest)
	if err != nil {
		s.logger.Err(err).Msg("create tournament request")
		return
	}

	g, err := getGender(createRequest.Gender)
	if err != nil {
		s.logger.Err(err).Msg("gender")
		response_factory.ReturnError(c, response_error.ValidationError)
		return
	}
	t := tournaments.Tournament{
		Alias:       createRequest.Alias,
		Name:        createRequest.Name,
		Description: createRequest.Description,
		Gender:      g,
	}

	ts, err := s.service.CreateTournament(c.Request.Context(), t)
	if err != nil {
		s.logger.Err(err).Msg("create tournament")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}
	response_factory.ReturnSuccess(c, response_success.FromTournamentResponse(ts))
}
