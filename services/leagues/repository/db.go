package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"vollyemsk_tournament_gateway/models/leagues"
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
