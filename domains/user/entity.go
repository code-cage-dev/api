package user

import "github.com/code-cage-dev/api/pkg/entity"

type Entity struct {
	entity.Base
	Username string `json:"username"`
	GithubID string `json:"github_id"`
}

type entityFields struct {
	ID       string
	GithubID string
}

var fields = entityFields{
	ID:       "id",
	GithubID: "github_id",
}
