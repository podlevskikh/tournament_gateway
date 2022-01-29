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

func (r *Db) GetPlayersByTeamAndGroup(ctx context.Context, teamID int, groupAlias string) ([]*groups.TeamGroupPlayer, error) {
	var ps []*groups.TeamGroupPlayer
	err := r.db.Set("_ctx", ctx).
		Preload("Player").
		Where("team_id = ?", teamID).
		Where("group_alias = ?", groupAlias).
		Find(&ps).Error
	if err != nil {
		return []*groups.TeamGroupPlayer{}, errors.Wrap(err, "get team")
	}
	return ps, nil
}
