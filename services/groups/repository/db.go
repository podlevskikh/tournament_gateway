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

func (r *Db) GetGroupsByTournamentSeasonStageLeagues(ctx context.Context, tournamentAlias, seasonAlias, stageAlias, leagueAlias string) ([]*groups.Group, error) {
	var grs []*groups.Group
	err := r.db.Set("_ctx", ctx).
		Preload("Tournament").
		Preload("Season").
		Preload("Stage").
		Preload("League").
		Where("tournament_alias = ?", tournamentAlias).
		Where("season_alias = ?", seasonAlias).
		Where("stage_alias = ?", stageAlias).
		Where("league_alias = ?", leagueAlias).
		Find(&grs).Error
	if err != nil {
		return []*groups.Group{}, errors.Wrap(err, "get groups")
	}
	return grs, nil
}

func (r *Db) GetGroupsByTournamentSeasonStages(ctx context.Context, tournamentAlias, seasonAlias, stageAlias string) ([]*groups.Group, error) {
	var grs []*groups.Group
	err := r.db.Set("_ctx", ctx).
		Preload("Tournament").
		Preload("Season").
		Preload("Stage").
		Preload("League").
		Where("tournament_alias = ?", tournamentAlias).
		Where("season_alias = ?", seasonAlias).
		Where("stage_alias = ?", stageAlias).
		Find(&grs).Error
	if err != nil {
		return []*groups.Group{}, errors.Wrap(err, "get groups")
	}
	return grs, nil
}
