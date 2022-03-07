package config

import (
	"os"

	"github.com/pedrocmart/canvas/internal/core/log"

	"github.com/pedrocmart/canvas/internal/core"
)

// NewConfig create new config
func NewConfig() (*core.Config, error) {
	env := os.Getenv("ENV")
	Host := os.Getenv("HOST")

	c := &core.Config{
		Log:  log.New(core.ServiceName, env, core.Version, os.Getenv("LOG_LEVEL")),
		Host: Host,
	}

	return c, nil
}
