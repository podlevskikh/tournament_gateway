package seasons

import (
	"github.com/pkg/errors"
	"time"
	"tournament_gateway/app/api/controllers/seasons/requests"
	"tournament_gateway/models/seasons"
)

func createStageFromRequest(seasonAlias string, stagesRequest []requests.StageRequest) ([]*seasons.Stage, error) {
	req := make([]*seasons.Stage, 0, len(stagesRequest))
	for _, s := range stagesRequest {
		t, err := getStageType(s.Type)
		if err != nil {
			return nil, errors.Wrap(err, "stage type")
		}
		req = append(req, &seasons.Stage{
			Alias:       s.Alias,
			ShortName:   s.ShortName,
			Name:        s.Name,
			DateStart:   time.Time(s.DateStart),
			DateFinish:  time.Time(s.DateFinish),
			IsCurrent:   s.IsCurrent,
			IconUrl:     s.IconUrl,
			SeasonAlias: seasonAlias,
			Type:        t,
		})
	}
	return req, nil
}

func getStageType(typeRequest string) (seasons.StageType, error) {
	switch typeRequest {
	case "groups":
		return seasons.StageTypeGroups, nil
	case "qualification":
		return seasons.StageTypeQualification, nil
	case "joint":
		return seasons.StageTypeJoint, nil
	case "playOff":
		return seasons.StageTypePlayOff, nil
	case "finalFour":
		return seasons.StageTypeFinalFour, nil
	}
	return seasons.StageTypeUndefined, errors.New("stage type undefined")
}
