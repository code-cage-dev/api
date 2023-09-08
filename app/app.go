package app

import (
	"github.com/code-cage-dev/api/app/command"
	"github.com/code-cage-dev/api/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	Login                command.LoginHandler
	ChallengeCreate      command.ChallengeCreateHandler
	ChallengeUpdate      command.ChallengeUpdateHandler
	ChallengeDelete      command.ChallengeDeleteHandler
	ChallengeMarkPrivate command.ChallengeMarkPrivateHandler
	ChallengeMarkPublic  command.ChallengeMarkPublicHandler
}

type Queries struct {
	CurrentUser   query.CurrentUserHandler
	ProfileView   query.ProfileViewHandler
	ChallengeView query.ChallengeViewHandler
}
