package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"tournament_gateway/models/groups"
)

type Db struct {
	db *gorm.DB
}

func NewDb(db *gorm.DB) *Db {
	return &Db{db: db}
}

func (r *Db) GetGroupsByTournamentSeasonStageLeagues(ctx context.Context, tournamentAlias, seasonAlias, stageAlias, leagueAlias string) ([]*groups.Group, error) {
	var grs []*groups.Group
	q := r.db.Set("_ctx", ctx).
		Preload("Tournament").
		Preload("Season").
		Preload("Stage").
		Preload("League")
	if tournamentAlias != "" {
		q = q.Where("tournament_alias = ?", tournamentAlias)
	}
	if seasonAlias != "" {
		q = q.Where("season_alias = ?", seasonAlias)
	}
	if stageAlias != "" {
		q = q.Where("stage_alias = ?", stageAlias)
	}
	if leagueAlias != "" {
		q = q.Where("league_alias = ?", leagueAlias)
	}
	err := q.Find(&grs).Error
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
		Preload("Matches.Gym").
		Preload("Matches.Gym.Metros").
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
		Preload("Matches.HomeTeam").
		Preload("Matches.GuestTeam").
		Preload("Matches.Result.SetResults").
		Preload("Matches.Gym").
		Preload("Matches.Gym.Metros").
		Where("alias = ?", groupAlias).
		First(&gr).Error
	if err != nil {
		return groups.Group{}, errors.Wrap(err, "get matches")
	}
	return gr, nil
}

func (r *Db) GetGroup(ctx context.Context, alias string) (*groups.Group, error) {
	var gr groups.Group
	err := r.db.Set("_ctx", ctx).
		Preload("Tournament").
		Preload("Season").
		Preload("Stage").
		Preload("League").
		Preload("GroupResults").
		Where("alias = ?", alias).
		First(&gr).Error
	if err != nil {
		return &groups.Group{}, errors.Wrap(err, "get group")
	}
	return &gr, nil
}

func (r *Db) UpdateGroup(ctx context.Context, g groups.Group) (*groups.Group, error) {
	err := r.db.Set("_ctx", ctx).Model(&g).Where("alias = ?", g.Alias).Update(&g).Error
	if err != nil {
		return nil, errors.Wrap(err, "save group")
	}
	return &g, nil
}

func (r *Db) CreateGroup(ctx context.Context, g groups.Group) (*groups.Group, error) {
	err := r.db.Set("_ctx", ctx).Create(&g).Error
	if err != nil {
		return nil, errors.Wrap(err, "save group")
	}
	return &g, nil
}

