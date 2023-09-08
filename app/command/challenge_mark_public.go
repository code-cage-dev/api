package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/code-cage-dev/api/domains/challenge"
	"github.com/google/uuid"
)

type ChallengeMarkPublicCmd struct {
	CurrentUserID uuid.UUID `json:"-"`
	ID            uuid.UUID `params:"id" validate:"required,uuid4"`
}

type ChallengeMarkPublicResult struct{}

type ChallengeMarkPublicHandler cqrs.Handler[ChallengeMarkPublicCmd, *ChallengeMarkPublicResult]

type challengeMarkPublicHandler struct {
	repo challenge.Repository
}

func NewChallengeMarkPublicHandler(repo challenge.Repository) ChallengeMarkPublicHandler {
	return &challengeMarkPublicHandler{repo: repo}
}

func (h *challengeMarkPublicHandler) Handle(ctx context.Context, cmd ChallengeMarkPublicCmd) (*ChallengeMarkPublicResult, *i18np.Error) {
	if err := h.repo.MarkPublic(ctx, cmd.ID, cmd.CurrentUserID); err != nil {
		return nil, err
	}
	return &ChallengeMarkPublicResult{}, nil
}
