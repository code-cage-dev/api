package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/code-cage-dev/api/clients/github"
)

type CurrentUserQuery struct {
	Token string `json:"-"`
}

type CurrentUserResult struct {
	User *github.User `json:"user"`
}

type CurrentUserHandler cqrs.Handler[CurrentUserQuery, *CurrentUserResult]

type currentUserHandler struct {
	client github.Client
}

func NewCurrentUserHandler(client github.Client) CurrentUserHandler {
	return &currentUserHandler{client: client}
}

func (h *currentUserHandler) Handle(ctx context.Context, query CurrentUserQuery) (*CurrentUserResult, *i18np.Error) {
	user, err := h.client.CurrentUser(ctx, query.Token)
	if err != nil {
		return nil, err
	}
	return &CurrentUserResult{
		User: user,
	}, nil
}
