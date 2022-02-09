package tournaments

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"tournament_gateway/app/api/controllers/tournaments/requests"
	"tournament_gateway/app/api/request_factory"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
	"tournament_gateway/models/tournaments"
)

type UpdateTournament struct {
	service Service
	logger  *zerolog.Logger
}

func NewUpdateTournament(service Service, logger *zerolog.Logger) *UpdateTournament {
	return &UpdateTournament{service: service, logger: logger}
}

func (s *UpdateTournament) HTTPHandler(c *gin.Context) {
	tournamentAlias := c.Param("tournament_alias")
	updateRequest := requests.UpdateTournamentRequest{}
	err := request_factory.ReadJSONRequestBody(c, &updateRequest)
	if err != nil {
		s.logger.Err(err).Msg("update tournament request")
		return
	}

	g, err := getGender(updateRequest.Gender)
	if err != nil {
		s.logger.Err(err).Msg("gender")
		response_factory.ReturnError(c, response_error.ValidationError)
		return
	}
	t, err := s.service.GetTournament(c.Request.Context(), tournamentAlias)
	if err != nil {
		s.logger.Err(err).Msg("get tournament") //todo 404
		response_factory.ReturnError(c, response_error.Internal)
		return
	}
	t.Name = updateRequest.Name
	t.Description = updateRequest.Description
	t.Gender = g

	ts, err := s.service.UpdateTournament(c.Request.Context(), *t)
	if err != nil {
		s.logger.Err(err).Msg("save tournament") //todo 404
		response_factory.ReturnError(c, response_error.Internal)
		return
	}
	response_factory.ReturnSuccess(c, response_success.FromTournamentResponse(ts))
}

func getGender(genderRequest string) (tournaments.TournamentGender, error) {
	switch genderRequest {
	case "male":
		return tournaments.TournamentGenderMale, nil
	case "female":
		return tournaments.TournamentGenderFemale, nil
	}
	return tournaments.TournamentGenderUndefined, errors.New("gender undefined")
}
