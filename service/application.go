package service

import (
	"github.com/code-cage-dev/api/app"
	"github.com/code-cage-dev/api/config"
)

type Config struct {
	App config.App
}

func NewApp(config Config) app.Application {
	return app.Application{
		Commands: app.Commands{},
		Queries:  app.Queries{},
	}
}
