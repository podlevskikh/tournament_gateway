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
