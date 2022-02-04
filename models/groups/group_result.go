package groups

const (
	WinsScoring   = "wins"
	PointsScoring = "points"
)

type GroupResult struct {
	ID          int
	GroupAlias  string
	Name        string
	ScoringType string
	Teams       []Team       `gorm:"many2many:team2group;foreignKey:ID;joinTableForeignKey:GroupResultID;associationForeignKey:ID;associationJoinTableForeignKey:TeamID"`
	TeamResults []TeamResult `gorm:"-"`
}

type TeamResult struct {
	Team                Team
	Place               int
	WinsScoringResult   *WinsScoringResult
	PointsScoringResult *PointsScoringResult
}

type WinsScoringResult struct {
	Wins               int
	MaxPossibleWins    int
	Points             int
	MaxPossiblePoints  int
	HandicapWins       int
	HandicapPoints     int
	ThreePointsMatches int
	TwoPointsMatches   int
	OnePointMatches    int
	ZeroPointsMatches  int
	WinSets            int
	LoseSets           int
	WinPoints          int
	LosePoints         int
}

type PointsScoringResult struct {
	Points             int
	MaxPossiblePoints  int
	HandicapPoints     int
	ThreePointsMatches int
	TwoPointsMatches   int
	OnePointMatches    int
	ZeroPointsMatches  int
	WinSets            int
	LoseSets           int
	WinPoints          int
	LosePoints         int
}
