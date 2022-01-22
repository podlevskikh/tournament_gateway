package seasons

import "time"

type Season struct {
	Alias      string
	Name       string
	DateStart  time.Time
	DateFinish time.Time
	IsCurrent  bool
}
