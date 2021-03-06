package seasons

import "time"

type StageType string

const (
	StageTypeUndefined     StageType = "undefined"
	StageTypeQualification StageType = "qualification"
	StageTypeGroups        StageType = "groups"
	StageTypeJoint         StageType = "joint"
	StageTypePlayOff       StageType = "playOff"
	StageTypeFinalFour     StageType = "finalFour"
)

type Stage struct {
	Alias       string
	ShortName   string
	Name        string
	DateStart   time.Time
	DateFinish  time.Time
	IsCurrent   bool
	IconUrl     string
	SeasonAlias string
	Type        StageType
}
