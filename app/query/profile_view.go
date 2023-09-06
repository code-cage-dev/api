package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/code-cage-dev/api/domains/user"
)

type ProfileViewQuery struct {
	UserName string `json:"user_name" validate:"required,min=1,max=255"`
}

type ProfileViewResult struct {
	User *user.Entity `json:"user"`
}

type ProfileViewHandler cqrs.Handler[ProfileViewQuery, *ProfileViewResult]

type profileViewHandler struct {
	repo user.Repository
}

func NewProfileViewHandler(repo user.Repository) ProfileViewHandler {
	return &profileViewHandler{repo: repo}
}

func (h *profileViewHandler) Handle(ctx context.Context, query ProfileViewQuery) (*ProfileViewResult, *i18np.Error) {
	user, err := h.repo.Get(ctx, query.UserName)
	if err != nil {
		return nil, err
	}
	return &ProfileViewResult{
		User: user,
	}, nil
}
