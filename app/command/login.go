package command

import (
	"context"
	"strconv"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/code-cage-dev/api/clients/github"
	"github.com/code-cage-dev/api/domains/user"
	"github.com/google/uuid"
)

type LoginCmd struct {
	Code  string `json:"code" validate:"required"`
	State string `json:"state" validate:"required"`
}

type LoginResult struct {
	AccessToken string       `json:"access_token"`
	User        *github.User `json:"user"`
	ID          uuid.UUID    `json:"id"`
}

type LoginHandler cqrs.Handler[LoginCmd, *LoginResult]

type loginHandler struct {
	repo   user.Repository
	client github.Client
}

func NewLoginHandler(repo user.Repository, client github.Client) LoginHandler {
	return &loginHandler{repo: repo, client: client}
}

func (h *loginHandler) Handle(ctx context.Context, cmd LoginCmd) (*LoginResult, *i18np.Error) {
	user, err := h.client.Access(ctx, &github.AccessRequest{
		Code:  cmd.Code,
		State: cmd.State,
	})
	if err != nil {
		return nil, err
	}
	res, err := h.repo.Login(ctx, strconv.Itoa(user.ID), user.Login)
	if err != nil {
		return nil, err
	}
	return &LoginResult{
		AccessToken: user.Token,
		User:        user,
		ID:          res.ID,
	}, nil
}
