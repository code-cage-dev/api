package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/code-cage-dev/api/domains/challenge"
	"github.com/google/uuid"
)

type ChallengeViewQuery struct {
	CurrentUserID uuid.UUID `json:"-"`
	Slug          string    `params:"slug" validate:"required,max=255,min=3"`
}

type ChallengeViewResult struct {
	ID                  uuid.UUID `json:"id"`
	Title               string    `json:"title"`
	Slug                string    `json:"slug"`
	Description         string    `json:"description"`
	IsPublic            bool      `json:"is_public"`
	DifficultyLevel     string    `json:"difficulty_level"`
	CreatedBy           uuid.UUID `json:"created_by"`
	PreferredLanguageID uuid.UUID `json:"preferred_language_id"`
}

type ChallengeViewHandler cqrs.Handler[ChallengeViewQuery, *ChallengeViewResult]

type challengeViewHandler struct {
	repo challenge.Repository
}

func NewChallengeViewHandler(repo challenge.Repository) ChallengeViewHandler {
	return &challengeViewHandler{repo: repo}
}

func (h *challengeViewHandler) Handle(ctx context.Context, query ChallengeViewQuery) (*ChallengeViewResult, *i18np.Error) {
	entity, err := h.repo.View(ctx, query.Slug, query.CurrentUserID)
	if err != nil {
		return nil, err
	}
	return &ChallengeViewResult{
		ID:                  entity.ID,
		Title:               entity.Title,
		Slug:                entity.Slug,
		Description:         entity.Description,
    IsPublic:           *entity.IsPublic,
		DifficultyLevel:     entity.DifficultyLevel.String(),
		CreatedBy:           entity.CreatedBy,
		PreferredLanguageID: entity.PreferredLanguageID,
	}, nil
}
