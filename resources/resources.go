package resources

import (
	"context"
	"github.com/jinzhu/gorm"
	"log"
)

type Resources struct {
	Env *Env
	DB  *gorm.DB
}

func Get(ctx context.Context, logger *log.Logger) *Resources {
	r := &Resources{}
	if err := r.initEnv(logger); err != nil {
		logger.Fatal(err)
	}
	if err := r.initDB(logger); err != nil {
		logger.Fatal(err)
	}
	return r
}
