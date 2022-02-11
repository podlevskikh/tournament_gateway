package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"tournament_gateway/models/leagues"
)

type Db struct {
	db *gorm.DB
}

func NewDb(db *gorm.DB) *Db {
	return &Db{db: db}
}

func (r *Db) GetLeagues(ctx context.Context) ([]*leagues.League, error) {
	var ls []*leagues.League
	err := r.db.Set("_ctx", ctx).Find(&ls).Error
	if err != nil {
		return []*leagues.League{}, errors.Wrap(err, "get leagues")
	}
	return ls, nil
}

func (r *Db) GetLeague(ctx context.Context, alias string) (*leagues.League, error) {
	var l leagues.League
	err := r.db.Set("_ctx", ctx).Where("alias = ?", alias).First(&l).Error
	if err != nil {
		return nil, errors.Wrap(err, "get league")
	}
	return &l, nil
}

func (r *Db) UpdateLeague(ctx context.Context, l leagues.League) (*leagues.League, error) {
	err := r.db.Set("_ctx", ctx).Model(&l).Where("alias = ?", l.Alias).Update(&l).Error
	if err != nil {
		return nil, errors.Wrap(err, "save league")
	}
	return &l, nil
}

func (r *Db) CreateLeague(ctx context.Context, l leagues.League) (*leagues.League, error) {
	err := r.db.Set("_ctx", ctx).Create(&l).Error
	if err != nil {
		return nil, errors.Wrap(err, "save league")
	}
	return &l, nil
}
