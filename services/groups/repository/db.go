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

func (r *Db) GetGroupWithTeams(ctx context.Context, groupAlias string) (groups.Group, error) {
	var gr groups.Group
	err := r.db.Set("_ctx", ctx).
		Preload("Teams").
		Where("alias = ?", groupAlias).
		First(&gr).Error
	if err != nil {
		return groups.Group{}, errors.Wrap(err, "get group with teams")
	}
	return gr, nil
}

func (r *Db) GetGroupWithGroupResultsAndMatches(ctx context.Context, groupAlias string) (groups.Group, error) {
	var gr groups.Group
	err := r.db.Set("_ctx", ctx).
		Preload("GroupResults").
		Preload("GroupResults.Teams").
		Preload("Matches").
		Preload("Matches.Result").
		Preload("Matches.Result.SetResults").
		Where("alias = ?", groupAlias).
		First(&gr).Error
	if err != nil {
		return groups.Group{}, errors.Wrap(err, "get group with teams")
	}
	return gr, nil
}

func (r *Db) GetGroupWithMatches(ctx context.Context, groupAlias string) (groups.Group, error) {
	var gr groups.Group
	err := r.db.Set("_ctx", ctx).
		Preload("Matches").
		Preload("Matches.Result").
		Preload("Matches.Result.SetResults").
		Where("alias = ?", groupAlias).
		First(&gr).Error
	if err != nil {
		return groups.Group{}, errors.Wrap(err, "get matches")
	}
	return gr, nil
}
