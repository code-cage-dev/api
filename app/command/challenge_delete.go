package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/code-cage-dev/api/domains/challenge"
	"github.com/google/uuid"
)

type ChallengeDeleteCmd struct {
	CurrentUserID uuid.UUID `json:"-"`
	ID            uuid.UUID `params:"id" validate:"required,uuid4"`
}

type ChallengeDeleteResult struct{}

type ChallengeDeleteHandler cqrs.Handler[ChallengeDeleteCmd, *ChallengeDeleteResult]

type challengeDeleteHandler struct {
	repo challenge.Repository
}

func NewChallengeDeleteHandler(repo challenge.Repository) ChallengeDeleteHandler {
	return &challengeDeleteHandler{repo: repo}
}

func (h *challengeDeleteHandler) Handle(ctx context.Context, cmd ChallengeDeleteCmd) (*ChallengeDeleteResult, *i18np.Error) {
	if err := h.repo.Delete(ctx, cmd.ID, cmd.CurrentUserID); err != nil {
		return nil, err
	}
	return &ChallengeDeleteResult{}, nil
}
