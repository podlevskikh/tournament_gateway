package response_success

import (
	"tournament_gateway/models/seasons"
)

type StagesResponse struct {
	Stages []StageResponse `json:"stages"`
}

type StageResponse struct {
	Alias       string `json:"alias"`
	ShortName   string `json:"shortName"`
	Name        string `json:"name"`
	DateStart   string `json:"dateStart"`
	DateFinish  string `json:"dateFinish"`
	IsCurrent   bool   `json:"isCurrent"`
	IconUrl     string `json:"iconUrl"`
	Type        string `json:"type"`
	SeasonAlias string `json:"seasonAlias"`
}

func FromStagesResponse(ss []*seasons.Stage) StagesResponse {
	res := make([]StageResponse, 0, len(ss))
	for _, s := range ss {
		res = append(res, FromStageResponse(s))
	}
	return StagesResponse{Stages: res}
}

func FromStageResponse(s *seasons.Stage) StageResponse {
	return StageResponse{
		Alias:       s.Alias,
		ShortName:   s.ShortName,
		Name:        s.Name,
		DateStart:   s.DateStart.Format("2006-01-02"),
		DateFinish:  s.DateFinish.Format("2006-01-02"),
		IsCurrent:   s.IsCurrent,
		IconUrl:     s.IconUrl,
		Type:        string(s.Type),
		SeasonAlias: s.SeasonAlias,
	}
}
