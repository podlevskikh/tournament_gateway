package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"tournament_gateway/models/seasons"
)

type Db struct {
	db *gorm.DB
}

func NewDb(db *gorm.DB) *Db {
	return &Db{db: db}
}

func (r *Db) GetSeasons(ctx context.Context) ([]*seasons.Season, error) {
	var s []*seasons.Season
	err := r.db.Set("_ctx", ctx).Find(&s).Error
	if err != nil {
		return []*seasons.Season{}, errors.Wrap(err, "get seasons")
	}
	return s, nil
}

func (r *Db) GetSeasonByAlias(ctx context.Context, alias string) (*seasons.Season, error) {
	var s seasons.Season
	err := r.db.Set("_ctx", ctx).Where("alias = ?", alias).Preload("Stages").First(&s).Error
	if err != nil {
		return nil, errors.Wrap(err, "get seasons")
	}
	return &s, nil
}

func (r *Db) UpdateStage(ctx context.Context, st *seasons.Stage) error {
	err := r.db.Set("_ctx", ctx).Model(&st).Where("alias = ?", st.Alias).Update(&st).Error
	if err != nil {
		return errors.Wrap(err, "update stage")
	}
	return nil
}

func (r *Db) CreateStage(ctx context.Context, st *seasons.Stage) error {
	err := r.db.Set("_ctx", ctx).Create(&st).Error
	if err != nil {
		return errors.Wrap(err, "create stage")
	}
	return nil
}

func (r *Db) UpdateSeason(ctx context.Context, se seasons.Season) error {
	se.Stages = []*seasons.Stage{}
	err := r.db.Set("_ctx", ctx).Model(&se).Where("alias = ?", se.Alias).Update(&se).Error
	if err != nil {
		return errors.Wrap(err, "update season")
	}
	return nil
}

func (r *Db) GetStageByAlias(ctx context.Context, alias string) (*seasons.Stage, error) {
	var s seasons.Stage
	err := r.db.Set("_ctx", ctx).Where("alias = ?", alias).First(&s).Error
	if err != nil {
		return nil, errors.Wrap(err, "ge stage")
	}
	return &s, nil
}

func (r *Db) CreateSeason(ctx context.Context, se seasons.Season) (*seasons.Season, error) {
	err := r.db.Set("_ctx", ctx).Create(&se).Error
	if err != nil {
		return nil, errors.Wrap(err, "create season")
	}
	return &se, nil
}
