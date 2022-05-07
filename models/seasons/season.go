package seasons

import "time"

type Season struct {
	Alias      string
	ShortName  string
	Name       string
	DateStart  time.Time
	DateFinish time.Time
	IsCurrent  bool
	Stages     []*Stage `gorm:"ForeignKey:season_alias;AssociationForeignKey:alias"`
}
