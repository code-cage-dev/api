package user

import "github.com/code-cage-dev/api/pkg/entity"

type Entity struct {
	entity.Base
	Username  string `json:"username"`
	GithubID  string `json:"github_id"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
}

type entityFields struct {
	ID       string
	GithubID string
	Username string
}

var fields = entityFields{
	ID:       "id",
	GithubID: "github_id",
	Username: "username",
}
