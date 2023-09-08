package service

import (
	"github.com/code-cage-dev/api/app"
	"github.com/code-cage-dev/api/app/command"
	"github.com/code-cage-dev/api/app/query"
	"github.com/code-cage-dev/api/clients/github"
	"github.com/code-cage-dev/api/config"
	"github.com/code-cage-dev/api/domains/challenge"
	"github.com/code-cage-dev/api/domains/user"
	"gorm.io/gorm"
)

type Config struct {
	App config.App
	DB  *gorm.DB
}

func NewApp(config Config) app.Application {
	githubClient := github.NewClient(config.App.Github)

	userRepo := user.NewRepo(config.DB)

	challengeRepo := challenge.NewRepo(config.DB)

	return app.Application{
		Commands: app.Commands{
			Login:           command.NewLoginHandler(userRepo, githubClient),
			ChallengeCreate: command.NewChallengeCreateHandler(challengeRepo),
			ChallengeUpdate: command.NewChallengeUpdateHandler(challengeRepo),
		},
		Queries: app.Queries{
			CurrentUser: query.NewCurrentUserHandler(githubClient),
			ProfileView: query.NewProfileViewHandler(userRepo),
		},
	}
}
