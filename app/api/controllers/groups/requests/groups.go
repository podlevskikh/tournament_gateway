package requests

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

type FilterRequest struct {
	TournamentAlias string `json:"tournamentAlias"`
	SeasonAlias     string `json:"seasonAlias"`
	StageAlias      string `json:"stageAlias"`
	LeagueAlias     string `json:"leagueAlias"`
}

func FilterRequestFromHTTPRequest(r *http.Request) (FilterRequest, error) {
	q := r.URL.Query()
	req := FilterRequest{}
	if q.Get("filter") != "" {
		err := json.Unmarshal([]byte(q.Get("filter")), &req)
		if err != nil {
			return FilterRequest{}, errors.Wrap(err, "filters not valid")
		}
	}
	return req, nil
}
