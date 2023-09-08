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

type ChallengeUpdateCmd struct {
	CurrentUserID       uuid.UUID `json:"-"`
	ID                  uuid.UUID `params:"id" validate:"required,uuid4"`
	Title               string    `json:"title" validate:"required,max=255,min=3"`
	PreferredLanguageID uuid.UUID `json:"preferred_language_id" validate:"required,uuid4"`
	Description         string    `json:"description" validate:"required,min=3,max=1000"`
	DifficultyLevel     string    `json:"difficulty_level" validate:"required,oneof=easy medium hard master"`
	IsPublic            *bool     `json:"is_public" validate:"required"`
}

type ChallengeUpdateResult struct {
	Entity *challenge.Entity `json:"challenge"`
}

type ChallengeUpdateHandler cqrs.Handler[ChallengeUpdateCmd, *ChallengeUpdateResult]

type challengeUpdateHandler struct {
	repo challenge.Repository
}

func NewChallengeUpdateHandler(repo challenge.Repository) ChallengeUpdateHandler {
	return &challengeUpdateHandler{repo: repo}
}

func (h *challengeUpdateHandler) Handle(ctx context.Context, cmd ChallengeUpdateCmd) (*ChallengeUpdateResult, *i18np.Error) {
	entity := &challenge.Entity{
		Title:               cmd.Title,
		Slug:                slug.New(cmd.Title),
		PreferredLanguageID: cmd.PreferredLanguageID,
		Description:         cmd.Description,
		DifficultyLevel:     challenge.DifficultyLevel(cmd.DifficultyLevel),
		IsPublic:            cmd.IsPublic,
		Base:                entity.UpdateBase(),
	}
	if err := h.repo.Update(ctx, cmd.ID, cmd.CurrentUserID, entity); err != nil {
		return nil, err
	}
	entity.CreatedBy = cmd.CurrentUserID
	return &ChallengeUpdateResult{
		Entity: entity,
	}, nil
}
