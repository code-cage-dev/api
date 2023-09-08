package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/code-cage-dev/api/domains/challenge"
	"github.com/google/uuid"
)

type ChallengeMarkPrivateCmd struct {
	CurrentUserID uuid.UUID `json:"-"`
	ID            uuid.UUID `params:"id" validate:"required,uuid4"`
}

type ChallengeMarkPrivateResult struct{}

type ChallengeMarkPrivateHandler cqrs.Handler[ChallengeMarkPrivateCmd, *ChallengeMarkPrivateResult]

type challengeMarkPrivateHandler struct {
	repo challenge.Repository
}

func NewChallengeMarkPrivateHandler(repo challenge.Repository) ChallengeMarkPrivateHandler {
	return &challengeMarkPrivateHandler{repo: repo}
}

func (h *challengeMarkPrivateHandler) Handle(ctx context.Context, cmd ChallengeMarkPrivateCmd) (*ChallengeMarkPrivateResult, *i18np.Error) {
	if err := h.repo.MarkPrivate(ctx, cmd.ID, cmd.CurrentUserID); err != nil {
		return nil, err
	}
	return &ChallengeMarkPrivateResult{}, nil
}
