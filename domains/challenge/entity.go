package challenge

import (
	"github.com/code-cage-dev/api/pkg/entity"
	"github.com/google/uuid"
)

type Entity struct {
	entity.Base
	CreatedBy         uuid.UUID `json:"created_by"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	DifficultyLevel   string    `json:"difficulty_level"`
	PreferredLanguage string    `json:"preferred_language"`
}
