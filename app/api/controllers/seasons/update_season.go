package seasons

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"time"
	"tournament_gateway/app/api/controllers/seasons/requests"
	"tournament_gateway/app/api/request_factory"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
	"tournament_gateway/app/api/response_success"
	"tournament_gateway/models/seasons"
)

type UpdateSeason struct {
	service Service
	logger  *zerolog.Logger
}

func NewUpdateSeason(service Service, logger *zerolog.Logger) *UpdateSeason {
	return &UpdateSeason{service: service, logger: logger}
}

func (s *UpdateSeason) HTTPHandler(c *gin.Context) {
	createRequest := requests.SeasonRequest{}
	err := request_factory.ReadJSONRequestBody(c, &createRequest)
	if err != nil {
		s.logger.Err(err).Msg("create season request")
		return
	}

	stages, err := createStageFromRequest(createRequest.Alias, createRequest.Stages)
	if err != nil {
		s.logger.Err(err).Msg("stage type")
		response_factory.ReturnError(c, response_error.ValidationError)
		return
	}
	se := seasons.Season{
		Alias:      createRequest.Alias,
		Name:       createRequest.Name,
		DateStart:  time.Time(createRequest.DateStart),
		DateFinish: time.Time(createRequest.DateFinish),
		IsCurrent:  createRequest.IsCurrent,
		Stages:     stages,
	}

	seas, err := s.service.UpdateSeason(c.Request.Context(), se)
	if err != nil {
		s.logger.Err(err).Msg("create season")
		response_factory.ReturnError(c, response_error.Internal)
		return
	}
	response_factory.ReturnSuccess(c, response_success.FromSeasonResponse(seas))
}
