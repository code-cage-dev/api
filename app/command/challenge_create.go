package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/code-cage-dev/api/domains/challenge"
	"github.com/code-cage-dev/api/pkg/entity"
	"github.com/google/uuid"
	"github.com/ssibrahimbas/slug"
)

type ChallengeCreateCmd struct {
	CurrentUserID       uuid.UUID `json:"-"`
	Title               string    `json:"title" validate:"required,max=255,min=3"`
	PreferredLanguageID uuid.UUID `json:"preferred_language_id" validate:"required,uuid4"`
	Description         string    `json:"description" validate:"required,min=3,max=1000"`
	DifficultyLevel     string    `json:"difficulty_level" validate:"required,oneof=easy medium hard master"`
	IsPublic            *bool     `json:"is_public" validate:"required"`
}

type ChallengeCreateResult struct {
	ID uuid.UUID `json:"id"`
}

type ChallengeCreateHandler cqrs.Handler[ChallengeCreateCmd, *ChallengeCreateResult]

type challengeCreateHandler struct {
	repo challenge.Repository
}

func NewChallengeCreateHandler(repo challenge.Repository) ChallengeCreateHandler {
	return &challengeCreateHandler{repo: repo}
}

func (h *challengeCreateHandler) Handle(ctx context.Context, cmd ChallengeCreateCmd) (*ChallengeCreateResult, *i18np.Error) {
	entity := &challenge.Entity{
		CreatedBy:           cmd.CurrentUserID,
		Title:               cmd.Title,
		Slug:                slug.New(cmd.Title),
		PreferredLanguageID: cmd.PreferredLanguageID,
		Description:         cmd.Description,
		DifficultyLevel:     challenge.DifficultyLevel(cmd.DifficultyLevel),
		IsPublic:            cmd.IsPublic,
		Base:                entity.DefaultBase(),
	}
	if err := h.repo.Create(ctx, entity); err != nil {
		return nil, err
	}
	return &ChallengeCreateResult{
		ID: entity.ID,
	}, nil
}
