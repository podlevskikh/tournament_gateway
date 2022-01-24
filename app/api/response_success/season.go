package response_success

import (
	"vollyemsk_tournament_gateway/models/seasons"
)

type SeasonsResponse struct {
	Seasons []SeasonResponse `json:"seasons"`
}

type SeasonResponse struct {
	Alias      string `json:"alias"`
	Name       string `json:"name"`
	DateStart  string `json:"dateStart"`
	DateFinish string `json:"dateFinish"`
	IsCurrent  bool   `json:"isCurrent"`
}

func FromSeasonResponse(ss []*seasons.Season) SeasonsResponse {
	res := make([]SeasonResponse, 0, len(ss))
	for _, s := range ss {
		res = append(res, SeasonResponse{
			Alias:      s.Alias,
			Name:       s.Name,
			DateStart:  s.DateStart.Format("2006-01-02"),
			DateFinish: s.DateFinish.Format("2006-01-02"),
			IsCurrent:  s.IsCurrent,
		})
	}
	return SeasonsResponse{Seasons: res}
}
