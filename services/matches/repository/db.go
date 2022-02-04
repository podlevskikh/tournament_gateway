package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"vollyemsk_tournament_gateway/models/groups"
)

type Db struct {
	db *gorm.DB
}

func NewDb(db *gorm.DB) *Db {
	return &Db{db: db}
}

func (r *Db) GetMatch(ctx context.Context, matchID int) (groups.Match, error) {
	var m groups.Match
	err := r.db.Set("_ctx", ctx).
		Preload("HomeTeam").
		Preload("GuestTeam").
		Preload("Result").
		Preload("Result.HomeTeamPlayers", "team = ?", groups.HomeTeam).
		Preload("Result.GuestTeamPlayers", "team = ?", groups.GuestTeam).
		Preload("Result.HomeBestPlayer").
		Preload("Result.GuestBestPlayer").
		Preload("Result.Referees").
		Preload("Result.SetResults").
		Where("ID = ?", matchID).
		Find(&m).Error
	if err != nil {
		return groups.Match{}, errors.Wrap(err, "get match")
	}
	return m, nil
}
