package requests

import (
	"strings"
	"time"
)

type SeasonRequest struct {
	Alias      string         `json:"alias"`
	Name       string         `json:"name"`
	DateStart  SeasonDate     `json:"dateStart"`
	DateFinish SeasonDate     `json:"dateFinish"`
	IsCurrent  bool           `json:"isCurrent"`
	Stages     []StageRequest `json:"stages"`
}

type StageRequest struct {
	Alias       string     `json:"alias"`
	Name        string     `json:"name"`
	DateStart   SeasonDate `json:"dateStart"`
	DateFinish  SeasonDate `json:"dateFinish"`
	IsCurrent   bool       `json:"isCurrent"`
	IconUrl     string     `json:"iconUrl"`
	Type        string     `json:"type"`
	SeasonAlias string     `json:"seasonAlias"`
}

type SeasonDate time.Time

func (j *SeasonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = SeasonDate(t)
	return nil
}
