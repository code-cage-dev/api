package user

import "github.com/code-cage-dev/api/pkg/entity"

type Entity struct {
	entity.Base
	Username string `json:"username"`
	Email    string `json:"email"`
	GithubID string `json:"github_id"`
}
