package response_success

import (
	"tournament_gateway/models/seasons"
)

type SeasonsResponse struct {
	Seasons []SeasonResponse `json:"seasons"`
}

type SeasonResponse struct {
	Alias      string          `json:"alias"`
	ShortName  string          `json:"shortName"`
	Name       string          `json:"name"`
	DateStart  string          `json:"dateStart"`
	DateFinish string          `json:"dateFinish"`
	IsCurrent  bool            `json:"isCurrent"`
	Stages     []StageResponse `json:"stages,omitempty"`
}

func FromSeasonsResponse(ss []*seasons.Season) SeasonsResponse {
	res := make([]SeasonResponse, 0, len(ss))
	for _, s := range ss {
		res = append(res, FromSeasonResponse(s))
	}
	return SeasonsResponse{Seasons: res}
}

func FromSeasonResponse(s *seasons.Season) SeasonResponse {
	return SeasonResponse{
		Alias:      s.Alias,
		ShortName:  s.ShortName,
		Name:       s.Name,
		DateStart:  s.DateStart.Format("2006-01-02"),
		DateFinish: s.DateFinish.Format("2006-01-02"),
		IsCurrent:  s.IsCurrent,
		Stages:     FromStagesResponse(s.Stages).Stages,
	}
}
