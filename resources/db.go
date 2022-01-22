package resources

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/pressly/goose"
	"log"
)

const driver = "mysql"

func (r *Resources) initDB(logger *log.Logger) error {
	conn, err := gorm.Open(driver, r.Env.DSN)
	if err != nil {
		return err
	}

	conn.SetLogger(logger)
	conn.LogMode(r.Env.Debug)

	r.DB = conn

	err = goose.SetDialect("mysql")
	if err != nil {
		return errors.Wrap(err, "set dialect")
	}
	err = goose.Up(r.DB.DB(), "./migrations")
	if err != nil {
		return errors.Wrap(err, "do migration")
	}
	logger.Println("done migration")

	logger.Println("db init")
	return nil
}
