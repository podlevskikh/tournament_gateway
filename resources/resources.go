package resources

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
)

type Resources struct {
	Env *Env
	DB  *gorm.DB
}

func Get(ctx context.Context, logger *zerolog.Logger) *Resources {
	r := &Resources{}
	if err := r.initEnv(logger); err != nil {
		logger.Fatal().Err(err).Msg("init ENV")
	}
	if err := r.initDB(logger); err != nil {
		logger.Fatal().Err(err).Msg("init DB")
	}
	return r
}
