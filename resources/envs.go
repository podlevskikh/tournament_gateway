package resources

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

type Env struct {
	ServiceName string `envconfig:"APP_SERVICENAME" default:"tournament_gateway"`
	PodName     string `envconfig:"APP_PODNAME" default:"podname"`

	DSN   string `envconfig:"DB_DSN" required:"true" example:"user:pass@tcp(host:port)/dbname"`
	Debug bool   `envconfig:"DB_DEBUG" default:"false"`
}

func init() {
	err := godotenv.Overload()
	if err != nil {
		fmt.Printf("failed to load .env file %s\n", err)
	}
}

func (r *Resources) initEnv(logger *zerolog.Logger) error {
	var s Env
	err := envconfig.Process("", &s)
	if err != nil {
		return err
	}

	r.Env = &s
	logger.Info().Msg("init env success")
	return nil
}
