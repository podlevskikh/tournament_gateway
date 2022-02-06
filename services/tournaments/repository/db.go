package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"tournament_gateway/models/tournaments"
)

type Db struct {
	db *gorm.DB
}

func NewDb(db *gorm.DB) *Db {
	return &Db{db: db}
}

func (r *Db) GetTournaments(ctx context.Context) ([]*tournaments.Tournament, error) {
	var t []*tournaments.Tournament
	err := r.db.Set("_ctx", ctx).Find(&t).Error
	if err != nil {
		return []*tournaments.Tournament{}, errors.Wrap(err, "get tournaments")
	}
	return t, nil
}