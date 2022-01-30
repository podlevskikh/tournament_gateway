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

func (r *Db) GetTeam(ctx context.Context, teamID int) (groups.Team, error) {
	var t groups.Team
	err := r.db.Set("_ctx", ctx).
		Preload("ContactUsers").
		Preload("HomeGyms").
		Preload("HomeGyms.Gym").
		Preload("HomeGyms.Gym.Metros").
		Where("id = ?", teamID).
		First(&t).Error
	if err != nil {
		return groups.Team{}, errors.Wrap(err, "get team")
	}
	return t, nil
}

func (r *Db) GetTeamWithGroups(ctx context.Context, teamID int) (groups.Team, error) {
	var t groups.Team
	err := r.db.Set("_ctx", ctx).
		Preload("Groups").
		Preload("Groups.Tournament").
		Preload("Groups.Season").
		Preload("Groups.Stage").
		Preload("Groups.League").
		Where("id = ?", teamID).
		First(&t).Error
	if err != nil {
		return groups.Team{}, errors.Wrap(err, "get team")
	}
	return t, nil
}
